package main

import (
	"bytes"
	"testing"
)

func TestGreetName(t *testing.T) {
	buffer := bytes.Buffer{}
	GreetName(&buffer, "Dom")

	got := buffer.String()
	want := "Hello, Dom!"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
