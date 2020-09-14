package model

import "github.com/awesomebusiness/uinvest/ent"

// ResponseAuthentication is response after success authentication
type ResponseAuthentication struct {
	*ent.User
	OTP int `json:"otp"`
}
