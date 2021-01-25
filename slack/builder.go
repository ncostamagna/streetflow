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
