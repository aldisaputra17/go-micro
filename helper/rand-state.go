package helper

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateRandomString(length int) (string, error) {
	byteLength := (length * 6) / 8

	randomBytes := make([]byte, byteLength)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	state := base64.URLEncoding.EncodeToString(randomBytes)

	if len(state) > length {
		state = state[:length]
	}

	return state, nil
}
