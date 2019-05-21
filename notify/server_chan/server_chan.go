package server_chan

import (
	"errors"
	"fmt"
	"github.com/GalvinGao/linkr/notify"
	"net/http"
	"net/url"
)

const (
	// ServerChan API endpoint; args: [sendToken]
	Api = "https://sc.ftqq.com/%s.send"
)

type Provider struct {
	ApiKey string
}

func New(apiKey string) Provider {
	return Provider{
		ApiKey: apiKey,
	}
}

// Send sends Message to serverchan
func (p Provider) Send(message notify.Message, _ notify.Extras) error {
	u := fmt.Sprintf(Api, p.ApiKey)

	response, err := http.PostForm(u, url.Values{
		"text": {message.Title},
		"desp": {message.Content},
	})
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("http status not ok")
	} else {
		return nil
	}
}
