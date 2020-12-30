package qpass_test

import (
	"math/rand"
	"qpass"
	"testing"
	"time"
)

func TestNewPassword(t *testing.T) {
	wantLen := 10
	g := qpass.NewGenerator()
	g.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	password := g.NewPassword(wantLen)
	gotLen := len(password)
	if wantLen != gotLen {
		t.Errorf("want %d, got %d", wantLen, gotLen)
	}
	letters := map[rune]struct{}{}
	for _, r := range password {
		letters[r] = struct{}{}
	}
	// This is a statistical test, which could still
	// pass by coincidence, but it's unlikely
	min := 4 // this feels about right
	if len(letters) < min {
		t.Errorf("want at least %d different letters in password: %q", min, password)
	}
}

func BenchmarkNewPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = qpass.NewPassword(10)
	}
}

func TestNewGenerator(t *testing.T) {
	wantLen := 10
	password := qpass.NewPassword(wantLen)
	gotLen := len(password)
	if wantLen != gotLen {
		t.Errorf("want %d, got %d", wantLen, gotLen)
	}
}