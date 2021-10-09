package main

import (
	"strings"
)

// Repeat returns the input string repeated N times
func Repeat(letter string, count int) string {

	var b strings.Builder
	for i := 1; i <= count; i++ {
		b.WriteString(letter)
	}

	return b.String()
}
