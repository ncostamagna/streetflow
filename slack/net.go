package slack

import (
	"github.com/slack-go/slack"
)

func (sb *SlackBuilder) doPostMessage(message string) (result *SlackResponse) {

	channelID, timestamp, err := sb.Client.PostMessage(
		sb.Channel,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)

	return &SlackResponse{
		Channel: channelID, Timestamp: timestamp, Err: err,
	}
}
