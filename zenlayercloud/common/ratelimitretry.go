package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	codeLimitExceeded = "REQUEST_LIMIT_EXCEEDED"
	rateLimitRetryFormat = "[WARN] rate limit exceeded, retrying (%d/%d) in %f seconds: %s"
)

func (c *Client) sendWithRateLimitRetry(request Request) (resp *http.Response, err error) {
	maxRetries := maxInt(c.config.RateLimitExceededMaxRetries, 0)
	if maxRetries == 0 {
		return c.sendWithNetworkFailureRetry(request)
	}

	durationFunc := safeDurationFunc(c.config.RateLimitExceededRetryDuration)
	ctx := request.GetContext()

	for idx := 0; idx <= maxRetries; idx++ {
		resp, err = c.sendWithNetworkFailureRetry(request)
		if err != nil {
			return
		}

		body, readErr := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if readErr != nil {
			return resp, readErr
		}
		resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		if idx < maxRetries && isRateLimitError(resp.StatusCode, body) {
			errResp := &ErrorResponse{}
			_ = json.Unmarshal(body, errResp)
			duration := durationFunc(idx)
			c.logger.Printf(rateLimitRetryFormat, idx, maxRetries, duration.Seconds(), errResp.Message)
			// Wait for the backoff duration, but bail out early if the
			// context is cancelled or its deadline is exceeded.
			select {
			case <-time.After(duration):
			case <-ctx.Done():
				return resp, ctx.Err()
			}
			continue
		}

		return resp, err
	}
	return resp, err
}

func isRateLimitError(statusCode int, body []byte) bool {
	if statusCode == http.StatusTooManyRequests {
		return true
	}
	errResp := &ErrorResponse{}
	if err := json.Unmarshal(body, errResp); err != nil {
		return false
	}
	return errResp.Code == codeLimitExceeded
}

func safeDurationFunc(durationFunc DurationFunc) DurationFunc {
	if durationFunc != nil {
		return durationFunc
	}
	return ExponentialBackoff
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
