package client

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"strings"
)

// Response ...
type Response struct {
	*http.Response
	Err      error
	byteBody []byte
}

// String return the Response Body as a String.
func (r *Response) String() string {
	return string(r.Bytes())
}

// Bytes return the Response Body as bytes.
func (r *Response) Bytes() []byte {
	return r.byteBody
}

func (r *Response) SetBytes(bytes []byte) {
	r.byteBody = bytes
}

// FillUp set the *fill* parameter with the corresponding JSON or XML response.
// fill could be `struct` or `map[string]interface{}`
func (r *Response) FillUp(fill interface{}) error {
	ctypeJSON := "application/json"
	ctypeXML := "application/xml"

	ctype := strings.ToLower(r.Header.Get("Content-Type"))

	for i := 0; i < 2; i++ {

		switch {
		case strings.Contains(ctype, ctypeJSON):
			return json.Unmarshal(r.byteBody, fill)
		case strings.Contains(ctype, ctypeXML):
			return xml.Unmarshal(r.byteBody, fill)
		case i == 0:
			ctype = http.DetectContentType(r.byteBody)
		}

	}

	return errors.New("response format neither JSON nor XML")
}
