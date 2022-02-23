package client

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func (rb *RequestBuilder) doRequest(verb string, reqURL string, reqBody interface{}) (result *Response) {

	start := time.Now()

	reqURL = rb.BaseURL + reqURL

	// Initialize response
	result = new(Response)

	body, err := rb.marshalReqBody(reqBody)
	if err != nil {
		result.Err = err
		if rb.LogTime {
			elapsed := time.Since(start)
			fmt.Printf("%s %s - time: %s\n", verb, reqURL, elapsed)
		}
		return result
	}

	mock := getMock(verb, reqURL)

	var httpResp *http.Response
	if mock != nil {
		httpResp = &http.Response{
			StatusCode: mock.RespHTTPCode,
			Body:       nopCloser{bytes.NewBufferString(mock.RespBody)},
		}
	} else {
		// Get Client (client + transport)
		client := rb.getClient()

		request, err := http.NewRequest(verb, reqURL, bytes.NewBuffer(body))
		if err != nil {
			result.Err = err
			if rb.LogTime {
				elapsed := time.Since(start)
				fmt.Printf("%s %s - time: %s\n", verb, reqURL, elapsed)
			}
			return result
		}

		// Set extra parameters
		rb.setParams(request)

		// Make the request
		httpResp, err = client.Do(request)
		if err != nil {
			result.Err = err
			if rb.LogTime {
				elapsed := time.Since(start)
				fmt.Printf("%s %s - time: %s\n", verb, reqURL, elapsed)
			}
			return result
		}
	}

	// Read response
	defer httpResp.Body.Close()
	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		result.Err = err
		if rb.LogTime {
			elapsed := time.Since(start)
			fmt.Printf("%s %s - time: %s\n", verb, reqURL, elapsed)
		}
		return result
	}

	result.Response = httpResp
	result.byteBody = respBody

	if rb.LogTime {
		elapsed := time.Since(start)
		fmt.Printf("%s %s - time: %s\n", verb, reqURL, elapsed)
	}
	return result
}

func (rb *RequestBuilder) marshalReqBody(body interface{}) (b []byte, err error) {
	if body != nil {
		switch rb.ContentType {
		case JSON:
			b, err = json.Marshal(body)
		case XML:
			b, err = xml.Marshal(body)
		case BYTES:
			var ok bool
			b, ok = body.([]byte)
			if !ok {
				err = fmt.Errorf("bytes: body is %T(%v) not a byte slice", body, body)
			}
		}
	}
	return
}

func (rb *RequestBuilder) getClient() *http.Client {

	defaultTransport := &http.Transport{
		DialContext: (&net.Dialer{Timeout: rb.getConnectionTimeout()}).DialContext,
	}

	tr := defaultTransport

	rb.Client = &http.Client{Transport: tr}

	return rb.Client
}

func (rb *RequestBuilder) getConnectionTimeout() time.Duration {

	switch {
	case rb.DisableTimeout:
		return 0
	case rb.ConnectTimeout > 0:
		return rb.ConnectTimeout
	default:
		return DefaultConnectTimeout
	}
}

func (rb *RequestBuilder) setParams(req *http.Request) {

	// Default headers
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "no-cache")

	// Encoding
	var cType string

	switch rb.ContentType {
	case JSON:
		cType = "json"
	case XML:
		cType = "xml"
	}

	if cType != "" {
		req.Header.Set("Accept", "application/"+cType)
		req.Header.Set("Content-Type", "application/"+cType)
	}

	for key, value := range rb.Headers {
		req.Header[key] = value
	}

}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

func getMock(verb, reqURL string) *Mock {

	if len(mockMap) < 1 {
		return nil
	}

	if mock := mockMap[verb+" "+reqURL]; mock != nil {
		return mock
	}

	return nil
}
