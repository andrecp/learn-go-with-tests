package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var countDownStart = 3
var finalWord = "Go!"

func Countdown(out io.Writer) error {
	for i := countDownStart; i >= 1; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)

	}

	time.Sleep(1 * time.Second)
	fmt.Fprint(out, "Go!")
	return nil
}

func main() {
	Countdown(os.Stdout)
}
