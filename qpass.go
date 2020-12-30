package qpass

import (
	"math/rand"
	"strings"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewPassword(length int) string {
	var s strings.Builder
	for i := 0; i < length; i++ {
		s.WriteRune('a' + rune(random.Intn(26)))
	}
	return s.String()
}
