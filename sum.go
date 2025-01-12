package main

func Sum(numbers []int) int {
	out := 0
	for _, number := range numbers {
		out += number
	}

	return out
}

func SumAll(slices ...[]int) []int {
	var output []int

	for _, slice := range slices {
		output = append(output, Sum(slice))
	}

	return output
}

func SumAllTails(slices ...[]int) []int {
	var output []int

	for _, slice := range slices {
		slice_sum := 0
		if len(slice) > 0 {
			slice_sum = Sum(slice[1:])
		}

		output = append(output, slice_sum)
	}

	return output
}
