package main

import (
	"fmt"
	"io"
)

func GreetName(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s!", name)
}
