package telegram

type TelegramBuilder struct {
	transport Transport
	telegram  *Telegram
}

func NewTelegramBuilder(transport Transport) *TelegramBuilder {
	return &TelegramBuilder{
		transport: transport,
		telegram: &Telegram{
			ChatID: chatID,
		},
	}
}

func (tb *TelegramBuilder) ChatID(chatID int64) *TelegramBuilder {
	tb.telegram.ChatID = chatID
	return tb
}

func (tb *TelegramBuilder) Message(text string) *TelegramBuilder {
	tb.telegram.Text = text
	return tb
}

func (tb *TelegramBuilder) Send() error {

	err := tb.transport.SendMesage(tb.telegram)
	return err
}
