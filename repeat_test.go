package main

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a", 5)
	expected := "aaaaa"

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
