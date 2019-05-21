package gotify

import (
	"encoding/json"
	"errors"
	"github.com/GalvinGao/linkr/notify"
	"net/http"
	"net/url"
	"path"
)

const (
	MessageEndpoint = "message"
)

type Provider struct {
	Endpoint       string
	ApplicationKey string
}

func New(endpoint string, appKey string) Provider {
	return Provider{
		Endpoint:       endpoint,
		ApplicationKey: appKey,
	}
}

func (p Provider) Send(message notify.Message, _ notify.Extras) error {
	u, err := url.Parse(p.Endpoint)
	if err != nil {
		return err
	}
	u.Path = path.Join(u.Path, MessageEndpoint)
	response, err := http.PostForm(u.String(), url.Values{
		"message": {message.Content},
		"title":   {message.Title},
	})
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		var jsonResult map[string]interface{}
		decoder := json.NewDecoder(response.Body)
		err = decoder.Decode(&jsonResult)
		if err != nil {
			return err
		}
		errorMessage := jsonResult["errorDescription"].(string)
		return errors.New(errorMessage)
	} else {
		return nil
	}
}
