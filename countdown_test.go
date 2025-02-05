package main

import (
	"bytes"
	"slices"
	"testing"
)

type SpyCountdown struct {
	Calls []string
}

func (s *SpyCountdown) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *SpyCountdown) Write([]byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return 0, nil
}

func TestCountdown(t *testing.T) {
	t.Run("prints countdown", func(t *testing.T) {
		spy_countdown := &SpyCountdown{}
		buffer := &bytes.Buffer{}
		Countdown(buffer, spy_countdown)

		got := buffer.String()
		want := "3\n2\n1\nGo!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("prints and sleeps in the right order", func(t *testing.T) {
		spy_countdown := &SpyCountdown{}
		Countdown(spy_countdown, spy_countdown)

		got := spy_countdown.Calls
		want := []string{"write", "sleep", "write", "sleep", "write", "sleep", "write"}

		if !slices.Equal(got, want) {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
