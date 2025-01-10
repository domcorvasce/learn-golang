package main

import "fmt"

const greetingPrefixEnglish = "Hello"
const greetingPrefixSpanish = "Hola"

func Hello() string {
	return greetingPrefixEnglish + ", world!"
}

func Greet(name string, lang string) string {
	if name == "" {
		return Hello()
	}

	prefix := greetingPrefixEnglish
	if lang == "Spanish" {
		prefix = greetingPrefixSpanish
	}

	return prefix + ", " + name + "!"
}

func main() {
	fmt.Println(Hello())
}
