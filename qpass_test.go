package qpass_test

import (
	"math/rand"
	"qpass"
	"testing"

	"github.com/google/go-cmp/cmp"
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
	g := qpass.NewGenerator(
		qpass.WithRandom(rand.New(rand.NewSource(1))),
	)
	want := "qTO$:(k l BD@G}[g/w]p5AK1E|sf/"
	got := g.NewPassword(15)
	if cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestUnicodeGenerator(t *testing.T) {
	g := qpass.NewGenerator(
		qpass.WithUnicode(),
		qpass.WithRandom(rand.New(rand.NewSource(1))),
	)
	want := "\U000b8b35\U00096f78\U000319db\U00087b0f\U000db7a5𨥌\U0008b029\U001043f7\U000f42fd\U00034b92\U0010968d\U0010c292\U000ee3ca\U000e2343\U000ad9f1"
	got := g.NewPassword(15)
	if cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func BenchmarkNewPassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = qpass.NewPassword(10)
	}
}
