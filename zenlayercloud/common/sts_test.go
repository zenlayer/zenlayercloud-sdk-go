package common

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"
)

const testSecurityToken = "zst_AbC-def_123"

type capturedRequest struct {
	header http.Header
	host   string
	body   string
}

func newStsServer(t *testing.T, captured *capturedRequest) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		captured.header = r.Header.Clone()
		captured.host = r.Host
		captured.body = string(body)
		_, _ = w.Write([]byte(`{"requestId":"r1","response":{"requestId":"r1"}}`))
	}))
}

func newStsClient(t *testing.T) *Client {
	c := &Client{}
	if err := c.InitWithStsCredential(NewStsCredential("STS.akid", "sts-secret", testSecurityToken)); err != nil {
		t.Fatal(err)
	}
	if err := c.WithConfig(NewConfig()); err != nil {
		t.Fatal(err)
	}
	return c
}

// gatewaySignature rebuilds the signature the way openapi-gateway does: signed header
// names and values are trimmed and lower-cased, then sorted by name.
func gatewaySignature(r *capturedRequest, signedHeaders []string, secret string) string {
	names := append([]string(nil), signedHeaders...)
	sort.Strings(names)
	var canonical strings.Builder
	for _, name := range names {
		value := r.header.Get(name)
		if name == "host" {
			value = r.host
		}
		canonical.WriteString(fmt.Sprintf("%s:%s\n", name, strings.ToLower(strings.TrimSpace(value))))
	}
	canonicalRequest := fmt.Sprintf("POST\n/\n\n%s\n%s\n%s", canonical.String(), strings.Join(signedHeaders, ";"), sha256hex(r.body))
	stringToSign := fmt.Sprintf("ZC2-HMAC-SHA256\n%s\n%s", r.header.Get("x-zc-timestamp"), sha256hex(canonicalRequest))
	return hex.EncodeToString([]byte(hmacsha256(stringToSign, secret)))
}

func TestStsCredentialSignsSecurityToken(t *testing.T) {
	captured := &capturedRequest{}
	server := newStsServer(t, captured)
	defer server.Close()

	_ = newStsClient(t).ApiCall(newTestRequest(server.URL), &BaseResponse{})

	if got := captured.header.Get(securityTokenHeader); got != testSecurityToken {
		t.Fatalf("expected security token header %q, got %q", testSecurityToken, got)
	}
	authorization := captured.header.Get("Authorization")
	if !strings.Contains(authorization, "Credential=STS.akid,") ||
		!strings.Contains(authorization, "SignedHeaders=content-type;host;x-zc-security-token,") {
		t.Fatalf("unexpected authorization %q", authorization)
	}
	signature := authorization[strings.Index(authorization, "Signature=")+len("Signature="):]
	expected := gatewaySignature(captured, []string{"content-type", "host", securityTokenHeader}, "sts-secret")
	if signature != expected {
		t.Fatalf("signature does not match the gateway's canonical form: got %s, want %s", signature, expected)
	}
}

func TestPlainCredentialDoesNotSendSecurityToken(t *testing.T) {
	captured := &capturedRequest{}
	server := newStsServer(t, captured)
	defer server.Close()

	c := &Client{}
	_ = c.InitWithCredential(NewCredential("akid", "secret"))
	_ = c.WithConfig(NewConfig())
	_ = c.ApiCall(newTestRequest(server.URL), &BaseResponse{})

	if _, ok := captured.header[http.CanonicalHeaderKey(securityTokenHeader)]; ok {
		t.Fatal("plain credential must not send a security token header")
	}
	if !strings.Contains(captured.header.Get("Authorization"), "SignedHeaders=content-type;host,") {
		t.Fatalf("unexpected authorization %q", captured.header.Get("Authorization"))
	}
}

func TestDebugDumpMasksSecurityToken(t *testing.T) {
	captured := &capturedRequest{}
	server := newStsServer(t, captured)
	defer server.Close()

	c := newStsClient(t)
	var logs bytes.Buffer
	c.debug = true
	c.logger = log.New(&logs, "", 0)
	_ = c.ApiCall(newTestRequest(server.URL), &BaseResponse{})

	if strings.Contains(logs.String(), testSecurityToken) {
		t.Fatal("debug output must not contain the security token")
	}
	if captured.header.Get(securityTokenHeader) != testSecurityToken {
		t.Fatal("masking the debug dump must not change the token sent on the wire")
	}
}

func TestInitWithStsCredentialRequiresAllFields(t *testing.T) {
	if err := (&Client{}).InitWithStsCredential(NewStsCredential("STS.akid", "sts-secret", "")); err == nil {
		t.Fatal("expected an error when the security token is missing")
	}
}
