package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(input string) bool {
	for _, char := range input {
		if char != '.' && char != '-' && char != ' ' {
			return false
		}
	}
	return true
}

func Convert(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("Empty line")
	}

	tabInput := strings.TrimSpace(input)
	if tabInput == "" {
		return "", fmt.Errorf("Only tab")
	}

	if isMorse(tabInput) {
		res := morse.ToText(tabInput)

		if res == "" {
			return "", fmt.Errorf("Error convert")
		}

		return res, nil
	}

	res := morse.ToMorse(tabInput)
	if res == "" {
		return "", fmt.Errorf("Error convert")
	}

	return res, nil
}
