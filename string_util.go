package main

import "strings"

func IntToString(input int) string {
	i := 1
	for input > i {
		i *= 10
	}
	i /= 10

	var b strings.Builder

	for i >= 1 {
		num := input / i % 10
		b.WriteRune(rune(num + 48))
		i /= 10
	}

	return b.String()
}
