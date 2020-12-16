package qpass

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewPassword(length int) string {
	var password []rune
	for i := 0; i < length; i++ {
		offset := random.Intn(26)
		password = append(password, rune('a' + offset))
	}
	return string(password)
}