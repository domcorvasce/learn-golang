package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
	want := "Hello, world!"

	assertCorrectMessage(t, got, want)
}

func TestGreet(t *testing.T) {
	t.Run("returns the provided name", func(t *testing.T) {
		got := Greet("Dom", "English")
		want := "Hello, Dom!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("defaults to 'Hello, world!' in case of empty name", func(t *testing.T) {
		got := Greet("", "English")
		want := "Hello, world!"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Greet("Elodie", "Spanish")
		want := "Hola, Elodie!"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
