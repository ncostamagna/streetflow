package slack

import (
	"github.com/slack-go/slack"
)

type SlackBuilder struct {
	Channel string
	Token   string
	Client  *slack.Client
}

func NewSlackBuilder(channel string, token string) *SlackBuilder {
	return &SlackBuilder{
		Channel: channel,
		Token:   token,
	}
}

func (sb *SlackBuilder) Build() (*SlackBuilder, error) {
	sb.Client = slack.New(sb.Token)

	return sb, nil
}

func (rb *SlackBuilder) SendMessage(message string) *SlackResponse {
	return rb.doPostMessage(message)
}

/*
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
} */
