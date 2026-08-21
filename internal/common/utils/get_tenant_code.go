package utils

import (
	"errors"
	"strings"
)

func GetTenantCode(input string) (string, error) {
	// Find the index of the '@' symbol
	parts := strings.Split(input, "@")

	// Ensure the input contains a valid '@' and has a part after it
	if len(parts) != 2 || parts[1] == "" {
		return "", errors.New("invalid format: tenant code not found")
	}

	return parts[1], nil
}
