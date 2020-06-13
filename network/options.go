package network

import (
	"net/http"
	"time"
)

type options struct {
	// Client
	jar     http.CookieJar
	timeout time.Duration

	// Transport
	keepAlive           time.Duration // default 30
	maxIdleConnsPerHost int           // default 2
	transport           *http.Transport
}

// Option the params of http which can self-defined
type Option func(*options)

// WithTimeout set the timeout of request
func WithTimeout(t time.Duration) Option {
	return func(o *options) {
		o.timeout = t
	}
}

// WithCookieJar set the CookieJar of request
func WithCookieJar(cj http.CookieJar) Option {
	return func(o *options) {
		o.jar = cj
	}
}

// WithTransport set the Transport of your own
func WithTransport(ts *http.Transport) Option {
	return func(o *options) {
		o.transport = ts
	}
}

// MaxIdleConnsPerHost set the max idle connects per host
func MaxIdleConnsPerHost(n int) Option {
	return func(o *options) {
		o.maxIdleConnsPerHost = n
	}
}

// KeepAlive set the connection keep live time
func KeepAlive(t time.Duration) Option {
	return func(o *options) {
		o.keepAlive = t
	}
}
