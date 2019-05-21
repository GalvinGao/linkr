package webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/GalvinGao/linkr/notify"
	"net/http"
)

type Provider struct {
	URL string
}

type Message struct {
	Action  string        `json:"action"`
	Content notify.Extras `json:"content"`
}

func New(url string) Provider {
	return Provider{
		URL: url,
	}
}

// Send sends Message to webhook
func (p Provider) Send(message notify.Message, extras notify.Extras) error {
	msg := Message{
		Action:  "new_visitor",
		Content: extras,
	}

	var jsonBuffer bytes.Buffer
	marshaled, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	jsonBuffer.Write(marshaled)

	response, err := http.Post(p.URL, "application/json", &jsonBuffer)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return errors.New("http status not ok")
	} else {
		return nil
	}
}
