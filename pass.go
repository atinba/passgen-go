package main

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	LOWER_CASE = "abcdefghijklmnopqrstuvwxyz"
	UPPER_CASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DIGITS     = "0123456789"
	SYMBOLS    = "!@#$%^&*()-_=+[]{}|;:,.<>?/`~"
)

func generate_password(length int) (string, error) {
	charset := LOWER_CASE + UPPER_CASE + DIGITS

	var password strings.Builder

	for i := 0; i < length; i++ {
		randIdx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}

		password.WriteByte(charset[randIdx.Int64()])
	}

	return password.String(), nil
}
