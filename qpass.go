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
	characterSet int
}

type Option func(*generator)

func WithUnicode() Option {
	return func(g *generator) {
		g.characterSet = Unicode
	}
}

func NewGenerator(opts ...Option) generator {
	g := generator{
		Rand:         rand.New(qrand.NewSource()),
		characterSet: Typeable,
	}
	for _, opt := range opts {
		opt(&g)
	}
	return g
}

func (g generator) Rune() rune {
	if g.characterSet == Typeable {
		return rune(' ' + g.Rand.Intn(95))
	}
	return rune(g.Rand.Intn(unicode.MaxRune + 1))
}

func (g generator) NewPassword(length int) string {
	var s strings.Builder
	for i := 0; i < length; i++ {
		s.WriteRune(g.Rune())
	}
	return s.String()
}
