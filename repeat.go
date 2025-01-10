package main

func Repeat(msg string, count int) string {
	output := ""

	for i := 0; i < count; i++ {
		output += msg
	}

	return output
}
