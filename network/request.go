package network

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

// Request wrap http request
type Request struct {
	cli      *Client
	method   string
	uri      string
	header   http.Header
	ctx      context.Context
	query    url.Values
	body     io.Reader
	respBody []byte
	err      error

	response *http.Response
}

func newRequest(cli *Client, uri string) *Request {
	return &Request{
		cli:    cli,
		uri:    uri,
		header: make(http.Header),
	}
}

// SetHeader set the http request header
func (r *Request) SetHeader(key, value string) *Request {
	r.header.Set(key, value)
	return r
}

// WithContext set the http request context
func (r *Request) WithContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

// SetBasicAuth  set the BasicAuth of the http request
func (r *Request) SetBasicAuth(username, password string) *Request {
	auth := username + ":" + password
	r.header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(auth)))
	return r
}

// Query add query parameter in the url
func (r *Request) Query(query url.Values) *Request {
	r.query = query
	return r
}

// Form append post from params to request
func (r *Request) Form(data url.Values) *Request {
	r.header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.body = strings.NewReader(data.Encode())
	return r
}

// JSON append post json params to request
func (r *Request) JSON(data interface{}) *Request {
	var err error
	r.header.Set("Content-Type", "application/json; charset=utf-8")
	buf := bytes.NewBuffer(nil)
	if data != nil {
		if err = json.NewEncoder(buf).Encode(data); err != nil {
			r.err = errors.Wrap(err, "request Encode")
		}
	}
	r.body = buf
	return r
}

// Body set put/post request body to request
func (r *Request) Body(in io.Reader) *Request {
	r.body = in
	return r
}

// ToJSON read response  of json format
func (r *Request) ToJSON(response interface{}) error {
	beanValue := reflect.ValueOf(response)
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("needs a pointer to a value")
	} else if beanValue.Elem().Kind() == reflect.Ptr {
		return errors.New("a pointer to a pointer is not allowed")
	}
	if err := r.execute(); err != nil {
		return err
	}

	// Unmarshal the response.
	if err := json.Unmarshal(r.respBody, response); err != nil {
		return errors.Wrap(err, "unmarshal response")
	}

	return nil
}

func (r *Request) Response() *http.Response {
	return r.response
}

// ToString read response  of string format
func (r *Request) ToString() (string, error) {
	if err := r.execute(); err != nil {
		return "", err
	}

	return string(r.respBody), nil
}

// ToBytes read response  of []byte format
func (r *Request) ToBytes() ([]byte, error) {
	if err := r.execute(); err != nil {
		return nil, err
	}

	return r.respBody, nil
}

func (r *Request) execute() error {
	if r.err != nil {
		return r.err
	}

	if r.query != nil {
		if !strings.Contains(r.uri, "?") {
			r.uri = r.uri + "?" + r.query.Encode()
		} else {
			r.uri = r.uri + r.query.Encode()
		}
	}
	request, err := http.NewRequest(r.method, r.uri, r.body)
	if err != nil {
		return errors.Wrap(err, "http new request")
	}
	request.Header = r.header
	if r.ctx != nil {
		request = request.WithContext(r.ctx)
	}

	//tracer
	traceReq, ht := TraceRequest(opentracing.GlobalTracer(), request)
	defer ht.Finish()
	//tracer

	resp, err := r.cli.Do(traceReq)

	if err != nil {
		return errors.Wrap(err, "http execute request")
	}

	r.response = resp

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read body")
	}

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		return fmt.Errorf("http response rrr: method: %s. status code: %d, body: %s",
			r.method, resp.StatusCode, b)
	}

	r.respBody = b

	return nil
}
