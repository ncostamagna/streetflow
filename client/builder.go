package client

import (
	"net/http"
	"time"
)

var DefaultTimeout = 500 * time.Millisecond

var DefaultConnectTimeout = 1500 * time.Millisecond

// DefaultMaxIdleConnsPerHost is the default maxium idle connections to have
// per Host for all clients, that use *any* RequestBuilder that don't set
// a CustomPool
var DefaultMaxIdleConnsPerHost = 2

// ContentType represents the Content Type for the Body of HTTP Verbs like
// POST, PUT, and PATCH
type ContentType int

const (
	// JSON represents a JSON Content Type
	JSON ContentType = iota

	// XML represents an XML Content Type
	XML

	// BYTES represents a plain Content Type
	BYTES
)

type RequestBuilder struct {

	// Headers to be send in the request
	Headers http.Header

	// Complete request time out.
	Timeout time.Duration

	//Connection timeout, it bounds the time spent obtaining a successful connection
	ConnectTimeout time.Duration

	// Base URL to be used for each Request. The final URL will be BaseURL + URL.
	BaseURL string

	// ContentType
	ContentType ContentType

	// Disable timeout and default timeout = no timeout
	DisableTimeout bool

	// Set an specific User Agent for this RequestBuilder
	UserAgent string

	// Public for custom fine tuning
	Client *http.Client

	LogTime bool
}

func (rb *RequestBuilder) Get(url string) *Response {
	return rb.doRequest(http.MethodGet, url, nil)
}

func (rb *RequestBuilder) Post(url string, body interface{}) *Response {
	return rb.doRequest(http.MethodPost, url, body)
}

func (rb *RequestBuilder) Put(url string, body interface{}) *Response {
	return rb.doRequest(http.MethodPut, url, body)
}

func (rb *RequestBuilder) Patch(url string, body interface{}) *Response {
	return rb.doRequest(http.MethodPatch, url, nil)
}

func (rb *RequestBuilder) Delete(url string) *Response {
	return rb.doRequest(http.MethodDelete, url, nil)
}

func (rb *RequestBuilder) Head(url string) *Response {
	return rb.doRequest(http.MethodHead, url, nil)
}

func (rb *RequestBuilder) Options(url string) *Response {
	return rb.doRequest(http.MethodOptions, url, nil)
}
