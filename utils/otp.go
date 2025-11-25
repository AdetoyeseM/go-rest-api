package utils

import (
	"crypto/rand"
	"fmt"
	"time"
)

func GenerateOTP() (string, time.Time) {
	// 6-digit numeric OTP
	b := make([]byte, 3)
	rand.Read(b)
	otp := fmt.Sprintf("%06d", int(b[0])<<16|int(b[1])<<8|int(b[2])%1000000)

	exp := time.Now().Add(10 * time.Minute)
	return otp, exp
}


