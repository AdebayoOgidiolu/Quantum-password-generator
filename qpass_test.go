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
	g := qpass.NewGenerator()
	g.Rand = rand.New(rand.NewSource(1))
	g.CharacterSet = qpass.Typeable
	wantPassword := "qTO$:(k l BD@G}[g/w]p5AK1E|sf/"
	password := g.NewPassword(30)
	if wantPassword != password {
		t.Errorf("want %q, got %q", wantPassword, password)
	}
	g.CharacterSet = qpass.Unicode
	wantPassword = "\U000b8b35\U00096f78\U000319db\U00087b0f\U000db7a5𨥌\U0008b029\U001043f7\U000f42fd\U00034b92\U0010968d\U0010c292\U000ee3ca\U000e2343\U000ad9f1"
	password = g.NewPassword(15)
	if wantPassword != password {
		t.Errorf("want %q, got %q", wantPassword, password)
	}
}

func BenchmarkNewPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = qpass.NewPassword(10)
	}
}
