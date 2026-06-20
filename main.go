package main

import (
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Not enough arguments!!!")
		return
	}

	word := os.Args[1]

	if _, err := ValidateInput(word); err != nil {
		fmt.Println("Input invalidated!!")
	}

	banner, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("Could not find banner file!!!")
		return
	}
	fmt.Println(Generate(word, banner))
}