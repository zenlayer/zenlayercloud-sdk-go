package common

import (
	"math"
	"net/http"
	"time"
)

// DurationFunc determines how long to wait between retries given the retry index.
type DurationFunc func(index int) time.Duration

// ExponentialBackoff returns exponentially increasing durations: 1s, 2s, 4s, 8s, ...
func ExponentialBackoff(index int) time.Duration {
	seconds := math.Pow(2, float64(index))
	return time.Duration(seconds) * time.Second
}

type Config struct {
	AutoRetry     bool `default:"false"`
	MaxRetryTime  int  `default:"3"`
	RetryDuration time.Duration
	HttpTransport *http.Transport   `default:""`
	Transport     http.RoundTripper `default:""`
	Proxy         string            `default:""`
	Scheme        string            `default:"HTTPS"`
	Domain        string            `default:""`
	// timeout in seconds
	Timeout int `default:"300"`
	Debug   *bool

	// RateLimitExceededMaxRetries defines the maximum number of retries when rate limit is exceeded.
	RateLimitExceededMaxRetries int `default:"3"`
	// RateLimitExceededRetryDuration determines the wait duration between rate limit retries.
	// Defaults to ExponentialBackoff if nil.
	RateLimitExceededRetryDuration DurationFunc
}

func NewConfig() (config *Config) {
	config = &Config{}
	InitStructWithDefaultTag(config)
	return
}

func (c *Config) WithAutoRetry(isAutoRetry bool) *Config {
	c.AutoRetry = isAutoRetry
	return c
}

func (c *Config) WithMaxRetryTime(maxRetryTime int) *Config {
	c.MaxRetryTime = maxRetryTime
	return c
}

func (c *Config) WithTimeout(timeout int) *Config {
	c.Timeout = timeout
	return c
}

func (c *Config) WithHttpTransport(httpTransport *http.Transport) *Config {
	c.HttpTransport = httpTransport
	return c
}

func (c *Config) WithScheme(scheme string) *Config {
	c.Scheme = scheme
	return c
}

func (c *Config) WithRateLimitRetry(maxRetries int, durationFunc DurationFunc) *Config {
	c.RateLimitExceededMaxRetries = maxRetries
	c.RateLimitExceededRetryDuration = durationFunc
	return c
}
