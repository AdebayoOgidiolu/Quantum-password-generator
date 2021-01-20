package qpass_test

import (
	"math/rand"
	"qpass"
	"testing"
)

func TestNewPassword(t *testing.T) {
	wantLen := 10
	password := qpass.NewPassword(wantLen)
	gotLen := len(password)
	if wantLen != gotLen {
		t.Errorf("want %d, got %d", wantLen, gotLen)
	}
}

func TestNewGenerator(t *testing.T) {
	wantLen := 30
	g := qpass.NewGenerator()
	g.Rand = rand.New(rand.NewSource(1))
	wantPassword := "qTO$:(k l BD@G}[g/w]p5AK1E|sf/"
	password := g.NewPassword(wantLen)
	if wantPassword != password {
		t.Errorf("want %q, got %q", wantPassword, password)
	}
}

func BenchmarkNewPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = qpass.NewPassword(10)
	}
}
