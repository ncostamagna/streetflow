package slack

type SlackResponse struct {
	Channel   string
	Timestamp string
	Err       error
	byteBody  []byte
}
