package main

func Sum(numbers []int) int {
	out := 0
	for _, number := range numbers {
		out += number
	}

	return out
}
