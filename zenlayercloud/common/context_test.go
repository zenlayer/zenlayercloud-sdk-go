package common

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

// newTestRequest builds a minimal POST request pointing at the given test server.
func newTestRequest(serverURL string) *BaseRequest {
	// httptest URL looks like "http://127.0.0.1:port"
	host := strings.TrimPrefix(serverURL, "http://")
	br := (&BaseRequest{}).Init()
	br.SetScheme("http")
	br.SetDomain(host)
	br.SetPath("/")
	br.SetBody([]byte("{}"))
	return br
}

func newTestClient() *Client {
	c := &Client{}
	_ = c.WithConfig(NewConfig())
	c.WithCredential(NewTokenCredential("test-token"))
	return c
}

// TestGetContextDefaultsToBackground verifies a request without an explicit
// context still yields a usable, non-nil context.
func TestGetContextDefaultsToBackground(t *testing.T) {
	br := (&BaseRequest{}).Init()
	if br.GetContext() != context.Background() {
		t.Fatalf("expected context.Background() by default, got %v", br.GetContext())
	}
}

// TestContextTimeoutAbortsRequest ensures a per-request timeout actually
// interrupts an in-flight HTTP call.
func TestContextTimeoutAbortsRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	req := newTestRequest(server.URL)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	req.SetContext(ctx)

	start := time.Now()
	_, err := newTestClient().sendWithNetworkFailureRetry(req)
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected an error due to context timeout, got nil")
	}
	if err != context.DeadlineExceeded {
		t.Fatalf("expected the raw context.DeadlineExceeded error, got %v", err)
	}
	if elapsed > 300*time.Millisecond {
		t.Fatalf("expected the call to abort quickly, took %s", elapsed)
	}
}

// TestContextCancellationBeforeCall ensures an already-cancelled context
// short-circuits before any HTTP request is made.
func TestContextCancellationBeforeCall(t *testing.T) {
	hit := false
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hit = true
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	req := newTestRequest(server.URL)
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately
	req.SetContext(ctx)

	_, err := newTestClient().sendWithNetworkFailureRetry(req)
	if err != context.Canceled {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
	if hit {
		t.Fatal("server should not have been called for an already-cancelled context")
	}
}
