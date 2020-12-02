package qpass

import (
	"math/rand"
	"time"
)

func NewPassword(length int) string {
	var random = rand.New(rand.NewSource(time.Now().UnixNano()))
	var password []byte
	var charSet = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for i := 0; i < length; i++ {
		index := random.Intn(len(charSet))
		password = append(password, charSet[index])
	}
	return string(password)
}