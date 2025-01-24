package main

import (
	"bytes"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep(duration time.Duration) {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	t.Run("prints countdown", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spy_sleeper := SpySleeper{}
		Countdown(buffer, &spy_sleeper)

		got := buffer.String()
		want := "3\n2\n1\nGo!\n"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if spy_sleeper.Calls != 3 {
			t.Fatal("expected sleep to get called 3 times")
		}
	})
}
