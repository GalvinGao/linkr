package main

import (
	"fmt"
	"github.com/GalvinGao/linkr/notify"
	"net/http"
)

type Notifier struct {
	Providers []notify.ServiceProvider
}

func NewNotifier(providers []notify.ServiceProvider) Notifier {
	return Notifier{
		Providers: providers,
	}
}

func (n Notifier) notify(req *http.Request, extras notify.Extras) chan error {
	userAgent := req.UserAgent()
	message := notify.Message{
		Title:   "New Visitor",
		Content: fmt.Sprintf("User-Agent: %s", userAgent),
	}

	extras.ClientHeaders = req.Header

	status := make(chan error)
	for _, provider := range n.Providers {
		go func() {
			status <- provider.Send(message, extras)
		}()
	}
	return status
}
