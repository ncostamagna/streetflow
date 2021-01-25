package telegram

const (
	url        = "https://api.telegram.org"
	chatID     = 1577933262
	tokenParam = "bot"
)

type Telegram struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}
