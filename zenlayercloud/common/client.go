package common

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"
)

// version this value will be replaced while build: -ldflags="-X 'common.version=x.x.x'"
var version = "0.2.44"

const SdkLang = "go"

type Client struct {
	credential    CredentialIface
	httpClient    *http.Client
	logger        *log.Logger
	config        *Config
	debug         bool
	requestClient string
}

// InitWithCredential initializes the client with HMAC-SHA256 signing credentials.
func (c *Client) InitWithCredential(credential *Credential) (err error) {
	if credential.SecretKeyId == "" || credential.SecretKeyPassword == "" {
		return NewZenlayerCloudSdkError(CredentialMissingError, "SecretKeyId or SecretKeyPassword is missing", "")
	}
	c.credential = credential
	return
}

// InitWithTokenCredential initializes the client with a Bearer token credential.
func (c *Client) InitWithTokenCredential(credential *TokenCredential) (err error) {
	if credential.Token == "" {
		return NewZenlayerCloudSdkError(CredentialMissingError, "Token is missing", "")
	}
	c.credential = credential
	return
}

// WithCredential sets the credential on the client and returns the client for chaining.
// Accepts any CredentialIface implementation (Credential or TokenCredential).
func (c *Client) WithCredential(credential CredentialIface) *Client {
	c.credential = credential
	return c
}

func (c *Client) WithConfig(config *Config) error {

	c.httpClient = &http.Client{}
	c.logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	c.config = config
	if config.Debug != nil {
		c.debug = *config.Debug
	} else {
		c.debug = os.Getenv("DEBUG") == "on"
	}
	if config.Transport != nil {
		c.httpClient.Transport = config.Transport
	} else if config.HttpTransport != nil {
		c.httpClient.Transport = config.HttpTransport
	}
	if config.Timeout > 0 {
		c.httpClient.Timeout = time.Duration(config.Timeout) * time.Second
	}

	// proxy & user agent
	if config.Proxy != "" {
		url, err := url.Parse(config.Proxy)
		if err != nil {
			return err
		}
		if c.httpClient.Transport == nil {
			c.logger.Printf("trying to set proxy when httpClient.Transport is nil")
		}
		if _, ok := c.httpClient.Transport.(*http.Transport); ok {
			c.httpClient.Transport.(*http.Transport).Proxy = http.ProxyURL(url)
		} else {
			c.logger.Println("setting proxy while httpClient.Transport is not a http.Transport is not supported")
		}
	}

	return nil
}

func (c *Client) WithRequestClient(rc string) *Client {
	const reRequestClient = "^[0-9a-zA-Z-_ ,;.]+$"

	if len(rc) > 128 {
		c.logger.Println("the length of RequestClient should be within 128 characters, it will be truncated")
		rc = rc[:128]
	}

	match, err := regexp.MatchString(reRequestClient, rc)
	if err != nil {
		c.logger.Println("regexp is wrong", reRequestClient)
		return c
	}
	if !match {
		c.logger.Printf("RequestClient not match the regexp: %s, ignored", reRequestClient)
		return c
	}
	c.requestClient = rc
	return c
}

func (c *Client) ApiCall(request Request, response Response) (err error) {
	if request.GetScheme() == "" {
		request.SetScheme(c.config.Scheme)
	}
	if request.GetDomain() == "" {
		request.SetDomain(c.config.Domain)
	}

	headers := request.GetHeaders()
	headers["x-zc-version"] = request.GetApiVersion()
	headers["x-zc-service"] = request.GetServiceName()
	headers["x-zc-action"] = request.GetAction()
	headers["x-zc-sdk-version"] = version
	headers["x-zc-sdk-lang"] = SdkLang
	headers["Host"] = request.GetDomain()
	if c.requestClient != "" {
		headers["x-zc-request-client"] = c.requestClient
	}

	request.SetContentType("application/json")

	body, err := json.Marshal(request)
	if err != nil {
		return err
	}
	request.SetBody(body)

	if token := c.credential.GetToken(); token != "" {
		headers["Authorization"] = "Bearer " + token
	} else {
		headers["x-zc-signature-method"] = "ZC2-HMAC-SHA256"
		authorization, err := c.signRequest(request)
		if err != nil {
			return err
		}
		headers["Authorization"] = authorization
	}

	httpResponse, err := c.apiCallWithRetry(request)
	if err != nil {
		return err
	}
	err = ParseFromHttpResponse(httpResponse, response)
	return err
}

func (c *Client) apiCallWithRetry(request Request) (resp *http.Response, err error) {
	return c.sendWithRateLimitRetry(request)
}

func (c *Client) sendWithNetworkFailureRetry(request Request) (resp *http.Response, err error) {
	var maxRetryTimes int
	var autoRetries bool
	if request.GetAutoRetries() {
		autoRetries = true
		maxRetryTimes = request.GetMaxAttempts()
	} else {
		autoRetries = c.config.AutoRetry
		maxRetryTimes = c.config.MaxRetryTime
	}

	ctx := request.GetContext()

	for retryTimes := 0; retryTimes <= maxRetryTimes; retryTimes++ {
		// Honor context cancellation/deadline before issuing each (re)try.
		if err = ctx.Err(); err != nil {
			return resp, err
		}

		httpRequest, err := http.NewRequest(request.GetHttpMethod(), request.GetUrl(), request.GetBodyReader())
		if err != nil {
			return nil, err
		}
		httpRequest = httpRequest.WithContext(ctx)

		for k, v := range request.GetHeaders() {
			httpRequest.Header[k] = []string{v}
		}

		resp, err = c.sendHttp(httpRequest)
		if err != nil && autoRetries && retryTimes < maxRetryTimes && ctx.Err() == nil {
			if netErr, ok := err.(net.Error); ok && (netErr.Timeout() || netErr.Temporary()) {
				duration := c.config.RetryDuration
				if c.debug {
					c.logger.Printf("[WARN] network failure, retrying (%d/%d) in %f seconds: %s", retryTimes, maxRetryTimes, duration.Seconds(), netErr.Error())
				}
				if duration > 0 {
					select {
					case <-time.After(duration):
					case <-ctx.Done():
						return resp, ctx.Err()
					}
				}
				continue
			}
		}
		if err != nil {
			// When the failure stems from the caller's context (cancellation
			// or deadline), surface the original context error so callers can
			// detect it via errors.Is(err, context.Canceled) etc., instead of
			// masking it as a generic network error.
			if ctxErr := ctx.Err(); ctxErr != nil {
				return resp, ctxErr
			}
			c.logger.Printf("[WARN] api request occurs error, action: %s, %v", request.GetAction(), err)
			msg := fmt.Sprintf("Fail to get response because %s", err)
			err = NewZenlayerCloudSdkError(NetworkError, msg, "")
		}
		return resp, err
	}
	return resp, err
}

func (c *Client) sendHttp(request *http.Request) (response *http.Response, err error) {
	if c.debug {
		outBytes, err := httputil.DumpRequest(request, true)
		if err != nil {
			c.logger.Printf("[ERROR] dump request failed, error %s", err)
			return nil, err
		}
		c.logger.Printf("[DEBUG] http request = %s", outBytes)
	}

	response, err = c.httpClient.Do(request)

	if c.debug && response != nil {
		outBytes, errRet := ioutil.ReadAll(response.Body)
		if errRet != nil {
			c.logger.Printf("[WAN] fail to read response body, err: %s", err.Error())
			return
		}
		response.Body = ioutil.NopCloser(bytes.NewBuffer(outBytes))
		c.logger.Printf("[DEBUG] http response = %s", string(outBytes))
	}
	return response, err
}

func (c *Client) signRequest(request Request) (string, error) {
	const algorithm = "ZC2-HMAC-SHA256"
	httpRequestMethod := request.GetHttpMethod()
	canonicalURI := "/"
	headers := request.GetHeaders()
	canonicalQueryString := ""

	hashedRequestPayload := sha256hex(request.GetBodyStr())
	canonicalHeaders := fmt.Sprintf("content-type:%s\nhost:%s\n", headers["Content-Type"], headers["Host"])
	signedHeaders := "content-type;host"

	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		httpRequestMethod,
		canonicalURI,
		canonicalQueryString,
		canonicalHeaders,
		signedHeaders,
		hashedRequestPayload)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	request.GetHeaders()["x-zc-timestamp"] = timestamp
	string2sign := fmt.Sprintf("%s\n%s\n%s",
		algorithm,
		timestamp,
		hashedCanonicalRequest)
	signature := hex.EncodeToString([]byte(hmacsha256(string2sign, c.credential.GetSecretKeyPassword())))
	authorization := fmt.Sprintf("%s Credential=%s, SignedHeaders=%s, Signature=%s",
		algorithm,
		c.credential.GetSecretKeyId(),
		signedHeaders,
		signature)
	return authorization, nil
}
