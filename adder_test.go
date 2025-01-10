package main

import "testing"

func TestAdd(t *testing.T) {
	t.Run("adds two integers", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertInt(t, sum, expected)
	})
}

func assertInt(t testing.TB, got, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}
