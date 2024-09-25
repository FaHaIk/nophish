package main

import (
	"fmt"
)

func ValidateServiceNumber(number string) bool {
	for index := range Services {
		numberGen := fmt.Sprintf("%02d", index)

		if numberGen == number {
			return true
		}
	}

	return false
}

func ValidateTemplateExists(number string) bool {
	for index := range Templates {
		numberGen := fmt.Sprintf("%02d", index)

		if numberGen == number {
			return true
		}
	}

	return false
}

func ValidateQorA(char string) bool {
	if char == "q" || char == "a" {
		return true
	}

	return false
}

func ValidateDomain(domain string) bool {
	return len(domain) > 0
}
