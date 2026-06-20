package main

import (
	"os"
	"strings"
	"fmt"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	input, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Missing file!!!")
		os.Exit(1)
	}

	lines := strings.Split(string(input), "\n")

	charMap := make(map[rune][]string)

	for i := 32; i < 127; i++ {
		start := (i-32)*9 + 1
		charMap[rune(i)] = lines[start : start+8]
	}

	return charMap, nil
}