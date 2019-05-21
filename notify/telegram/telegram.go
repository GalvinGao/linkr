package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GalvinGao/linkr/notify"
	"net/http"
	"net/url"
	"text/template"
	"unsafe"
)

const (
	// Telegram API endpoint; args: [botToken, methodName]
	Api = "https://api.telegram.org/bot%s/%s"
	// Message Template; Supporting Markdown: true
	MessageTmpl = `**{{.Title}}**
{{.Content}}`
)

type Provider struct {
	botToken string
	chatId   uint32
}

func New(botToken string, chatId uint32) Provider {
	return Provider{
		botToken: botToken,
		chatId:   chatId,
	}
}

// Send sends Message to telegram
func (p Provider) Send(message notify.Message, _ notify.Extras) error {
	t := template.Must(template.New("message").Parse(MessageTmpl))
	u := fmt.Sprintf(Api, p.botToken, "sendMessage")
	var writer bytes.Buffer
	if err := t.Execute(&writer, message); err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(&writer); err != nil {
		return err
	}
	b := buf.Bytes()
	s := *(*string)(unsafe.Pointer(&b))

	response, err := http.PostForm(u, url.Values{
		"chat_id":    {string(p.chatId)},
		"text":       {s},
		"parse_mode": {"Markdown"},
	})
	if err != nil {
		return err
	}

	var jsonResult map[string]interface{}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&jsonResult)
	if err != nil {
		return err
	}
	success := jsonResult["ok"].(bool)
	if success {
		return nil
	} else {
		return errors.New("telegram responds with a not-ok signal")
	}
}
