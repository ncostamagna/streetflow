package telegram

import (
	"fmt"
	"time"

	c "github.com/ncostamagna/streetflow/client"
)

// Transport -
type Transport interface {
	SendMesage(telegram *Telegram) error
	BaseURL(url string) Transport
}

type clientHTTP struct {
	token  string
	client c.RequestBuilder
}

// ClientType -
type ClientType int

const (
	// HTTP transport type
	HTTP ClientType = iota

	// Socket transport type
	Socket

	// GRPC transport type
	GRPC
)

// NewClient -
func NewClient(token string, ct ClientType) Transport {

	switch ct {
	case HTTP:
		return &clientHTTP{
			token: token,
			client: c.RequestBuilder{
				BaseURL:        fmt.Sprintf("%s/%s%s/", url, tokenParam, token),
				ConnectTimeout: 5000 * time.Millisecond,
				LogTime:        true,
			},
		}
	}

	panic("not a valid client")

}

func (c *clientHTTP) BaseURL(url string) Transport {
	c.client.BaseURL = fmt.Sprintf("%s/%s%s/", url, tokenParam, c.token)
	return c
}

func (c *clientHTTP) SendMesage(telegram *Telegram) error {

	reps := c.client.Post("sendMessage", telegram)

	if reps.Err != nil {
		return reps.Err
	}

	if reps.StatusCode > 299 {
		return fmt.Errorf("code: %d, message: %s", reps.StatusCode, reps)
	}

	/* 	if err := json.Unmarshal(reps.Bytes(), survey); err != nil {
		return err
	} */

	return nil

}
