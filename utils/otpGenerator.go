package utils

import (
	"crypto/rand"
	"sentinel/logger"
)

func GenerateOTP(length int) (string, error) {
	const otpChars = "1234567890"

	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		logger.LogMessageInRed("Not able to generate OTP code, cannot read buffer: " + err.Error())
		return "", err
	}
	otpCharsLength := len(otpChars)

	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}
