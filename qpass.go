package qpass

import (
	"math/rand"
	"strings"

	"github.com/bitfield/qrand"
)

func NewPassword(length int) string {
	return NewGenerator().NewPassword(length)
}

type generator struct{
	Rand *rand.Rand
}

func NewGenerator() generator {
	return generator{
		Rand: rand.New(qrand.NewSource()),
	}
}

func (g generator) NewPassword(length int) string {
	var s strings.Builder
	for i := 0; i < length; i++ {
		s.WriteRune('a' + rune(g.Rand.Intn(26)))
	}
	return s.String()
}