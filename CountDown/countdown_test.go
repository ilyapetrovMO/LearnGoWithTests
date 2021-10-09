package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpySleeper struct {
	Calls []string
}

func (s *SpySleeper) Sleep(amount time.Duration) {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpySleeper) Write(b []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("print 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpySleeper{})

		want := "3\n2\n1\nGo!"

		if buffer.String() != want {
			t.Errorf("got %q want %q", buffer.String(), want)
		}
	})

	t.Run("Sleep before every print", func(t *testing.T) {
		spySleeper := &SpySleeper{}
		Countdown(spySleeper, spySleeper)

		want := []string{sleep, write, sleep, write, sleep, write, sleep, write}

		if !reflect.DeepEqual(want, spySleeper.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeper.Calls)
		}
	})
}
