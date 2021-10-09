package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type DefaultSleeper struct{}

func (d DefaultSleeper) Sleep(t time.Duration) {
	time.Sleep(t)
}

type Sleeper interface {
	Sleep(time.Duration)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i >= 1; i-- {
		s.Sleep(1 * time.Second)
		fmt.Fprintln(w, i)
	}

	s.Sleep(1 * time.Second)
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
