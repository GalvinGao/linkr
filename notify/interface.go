package notify

import "net/http"

type Message struct {
	Title   string
	Content string
}

type Extras struct {
	ClientHeaders http.Header `json:"client_headers"`
	ShortUrl      string      `json:"short_url"`
	LongUrl       string      `json:"long_url"`
}

type ServiceProvider interface {
	Send(message Message, extras Extras) error
}
