package qpass

import (
	"math/rand"
	"strings"
	"unicode"

	"github.com/bitfield/qrand"
)

const (
	Typeable = iota
	Unicode
)

func NewPassword(length int) string {
	return NewGenerator().NewPassword(length)
}

type generator struct {
	Rand         *rand.Rand
	CharacterSet int
}

func NewGenerator() generator {
	return generator{
		Rand:         rand.New(qrand.NewSource()),
		CharacterSet: Typeable,
	}
}

func (g generator) typeableRune() rune {
	return rune(' ' + g.Rand.Intn(95))
}

func (g generator) unicodeRune() rune {
	return rune(g.Rand.Intn(unicode.MaxRune + 1))
}

func (g generator) NewPassword(length int) string {
	var s strings.Builder
	var newRune func() rune
	if g.CharacterSet == Unicode {
		newRune = g.unicodeRune
	} else {
		newRune = g.typeableRune
	}
	for i := 0; i < length; i++ {
		s.WriteRune(newRune())
	}
	return s.String()
}
