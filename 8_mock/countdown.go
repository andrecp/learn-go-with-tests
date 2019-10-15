package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var countDownStart = 3
var finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(out io.Writer, sleep Sleeper) error {
	for i := countDownStart; i >= 1; i-- {
		fmt.Fprintln(out, i)
		sleep.Sleep()

	}

	sleep.Sleep()
	fmt.Fprint(out, "Go!")
	return nil
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
