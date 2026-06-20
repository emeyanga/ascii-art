package main

import (
	"errors"
)

func ValidateInput(s string) (string, error) {
	for _, ch := range s {
		if ch < 32 || ch > 126 {
			return string(ch), errors.New("Input is invalid")
		}
	}
	return s, nil
}