package twillio

import "net/http"

// TwillioMessage is hold all method that will be implementation in twillio client
type TwillioMessage interface {
	SendOTP(toUserNumber string, name string) (*http.Response, int, error)
}
