package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Split(s, " ")
	m     := make(map[string] int)

	for _, value := range words {
		m[value]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
