package util

import (
	"regexp"
)

// IsEmailValid is validation to check email valid or not and return boolean
func IsEmailValid(email string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if email == "" || !emailRegex.MatchString(email) {
		return false
	}

	return true
}

// IsPhoneNumberValid is validation to check whether phone number valid or not and return boolean
func IsPhoneNumberValid(phonenumber string) bool {
	if phonenumber == "" {
		return false
	}

	phonenumberChar := string([]rune(phonenumber)[0]) + string([]rune(phonenumber)[1]) + string([]rune(phonenumber)[2])

	if len(phonenumber) < 12 || phonenumberChar != "+62" {
		return false
	}

	return true
}
