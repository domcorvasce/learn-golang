package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("sums slice of integers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})

	t.Run("sums empty slice of integers", func(t *testing.T) {
		numbers := []int{}

		got := Sum(numbers)
		expected := 0

		if got != expected {
			t.Errorf("expected %d, got %d, given %v", expected, got, numbers)
		}
	})
}
