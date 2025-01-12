package main

import (
	"slices"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	t.Run("sums two slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !slices.Equal(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sums two slice tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 9})
		expected := []int{2, 9}

		if !slices.Equal(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		if !slices.Equal(got, expected) {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})
}
