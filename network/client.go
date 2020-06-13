package network

import (
	"net"
	"net/http"
	"time"
)

// Client the http client wrap
type Client struct {
	*http.Client
}

// NewClient return a http client wrap to deal with http request
func NewClient(opt ...Option) *Client {
	opts := options{
		timeout:             5 * time.Second, // 请求超时时间
		keepAlive:           30 * time.Second,
		maxIdleConnsPerHost: 2, // 请求量较大时需调整此参数,否则会出现fd被耗尽,出现大量TIME_WAIT
	}
	for _, o := range opt {
		o(&opts)
	}
	var ts *http.Transport
	if opts.transport != nil {
		ts = opts.transport
	} else {
		ts = &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: opts.keepAlive,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   opts.maxIdleConnsPerHost,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}
	}

	return &Client{
		&http.Client{
			Timeout: opts.timeout,
			Jar:     opts.jar,
			Transport: &Transport{
				RoundTripper: ts,
			},
		}}
}

// Head create a new http head request
func (c *Client) Head(uri string) *Request {
	r := newRequest(c, uri)
	r.method = "HEAD"
	return r
}

// Get create a new http get request
func (c *Client) Get(uri string) *Request {
	r := newRequest(c, uri)
	r.method = "GET"
	return r
}

// Post create a new http post request
func (c *Client) Post(uri string) *Request {
	r := newRequest(c, uri)
	r.method = "POST"
	return r
}

// Put create a new http put request
func (c *Client) Put(uri string) *Request {
	r := newRequest(c, uri)
	r.method = "PUT"
	return r
}

// Delete create a new http delete request
func (c *Client) Delete(uri string) *Request {
	r := newRequest(c, uri)
	r.method = "DELETE"
	return r
}
