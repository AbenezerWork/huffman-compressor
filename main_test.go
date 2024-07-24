package main

import (
	"testing"
)

func TestCount(t *testing.T) {
	got := Count('t')
	want := int64(223_000)
	if got != want {
		t.Errorf("got %v want % v", got, want)
	}
	got = Count('X')
	want = int64(333)
	if got != want {
		t.Errorf("got %v want % v", got, want)
	}
}
