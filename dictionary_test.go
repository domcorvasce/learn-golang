package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("lookup hit", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertLookup(t, got, want, "test")
	})

	t.Run("lookup miss", func(t *testing.T) {
		_, err := dictionary.Search("foobar")
		want := "cannot retrieve value for key 'foobar'"

		if err == nil {
			t.Fatal("expected an error on lookup miss")
		}

		assertLookup(t, err.Error(), want, "foobar")
	})
}

func TestDictAdd(t *testing.T) {
	t.Run("term did not previously exist", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		got, err := dictionary.Search("test")
		if err != nil {
			t.Fatal("expected to found value for key")
		}

		assertLookup(t, got, "this is just a test", "test")
	})

	t.Run("term previously existed", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("test", "this is just a test")

		err := dictionary.Add("test", "this is just another test")
		if err == nil {
			t.Fatal("expected error on overwriting existing term")
		}

		assertLookup(t, err.Error(), "cannot overwrite existing term", "test")
	})
}

func assertLookup(t testing.TB, got, want, given string) {
	if got != want {
		t.Errorf("got %q, want %q, given %q", got, want, given)
	}
}
