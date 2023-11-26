package utils

import "crypto/rand"

func NewID() string {
	b := make([]byte, 20)
	rand.Read(b)
	return string(b)
}
