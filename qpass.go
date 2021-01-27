package qpass

import (
	"math/rand"
	"strings"
	"unicode"

	"github.com/bitfield/qrand"
)

/*const (
	Typeable = iota
	Unicode
)*/

var characterSet string

func NewPassword(length int, characterSet string) string {
	return NewGenerator().NewPassword(length, characterSet)
}

type generator struct {
	Rand         *rand.Rand
	CharacterSet string
}

func NewGenerator() generator {
	return generator{
		Rand:         rand.New(qrand.NewSource()),
		CharacterSet: "",
	}
}

func (g generator) typeableRune() rune {
	return rune(' ' + g.Rand.Intn(95))
}

func (g generator) unicodeRune() rune {
	return rune(g.Rand.Intn(unicode.MaxRune + 1))
}

func (g generator) NewPassword(length int, characterSet string) string {
	var s strings.Builder
	var newRune func() rune
	if g.CharacterSet == "unicode" {
		newRune = g.unicodeRune
	} else {
		newRune = g.typeableRune
	}
	for i := 0; i < length; i++ {
		s.WriteRune(newRune())
	}
	return s.String()
}
