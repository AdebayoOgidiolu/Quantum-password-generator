package qpass_test

import (
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