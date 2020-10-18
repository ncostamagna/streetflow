package rest

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func (rb *RequestBuilder) doRequest(verb string, reqURL string, reqBody interface{}) (result *Response) {

	fmt.Println(rb.BaseURL)
	reqURL = rb.BaseURL + reqURL

	result = new(Response)

	body, err := rb.marshalReqBody(reqBody)
	if err != nil {

		result.Err = err
		return
	}

	//Get Client (client + transport)
	client := rb.getClient()

	request, err := http.NewRequest(verb, reqURL, bytes.NewBuffer(body))
	if err != nil {

		result.Err = err
		return result
	}

	// Set extra parameters
	rb.setParams(request)

	// Make the request
	httpResp, err := client.Do(request)
	if err != nil {
		result.Err = err
		return result
	}

	// Read response
	defer httpResp.Body.Close()
	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		result.Err = err
		return result
	}

	result.Response = httpResp
	result.byteBody = respBody

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

	fmt.Println("client 1")
	defaultTransport := &http.Transport{
		//MaxIdleConnsPerHost:   DefaultMaxIdleConnsPerHost,
		//Proxy:                 http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{Timeout: rb.getConnectionTimeout()}).DialContext,
		//ResponseHeaderTimeout: rb.getRequestTimeout(),
	}

	tr := defaultTransport

	rb.Client = &http.Client{Transport: tr}

	return rb.Client
}

/* func (rb *RequestBuilder) getRequestTimeout() time.Duration {

	switch {
	case rb.DisableTimeout:
		return 0
	case rb.Timeout > 0:
		return rb.Timeout
	default:
		return DefaultTimeout
	}
} */

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

	//Default headers
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cache-Control", "no-cache")

	//Encoding
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

}
