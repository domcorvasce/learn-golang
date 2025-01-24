package main

import (
	"fmt"
	"io"
	"time"
)

type Sleeper interface {
	Sleep(duration time.Duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep(1 + time.Second)
	}

	fmt.Fprintln(out, "Go!")
}
