package pkg

import "net/http"

// TwillioMessage is hold all method that will be implementation in twillio client
type TwillioMessage interface {
	SendMessage(toUserNumber string, message string) (*http.Response, error)
}
