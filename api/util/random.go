package util

import (
	"math/rand"
	"time"
)

// GenerateRandomOTP is generate random number for otp and return integer
func GenerateRandomOTP() int {
	rand.Seed(time.Now().UnixNano())

	randomOTP := rand.Intn(2000-1000+1) + 1000

	return randomOTP
}
