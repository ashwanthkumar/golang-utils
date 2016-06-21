package io

import (
	"bufio"
	"os"
)

// ReadFullyFromStdin reads the STDIN and return it as a string
func ReadFullyFromStdin() string {
	input := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		input = input + line
	}

	return input
}

// ReadFullyFromFile reads the entire file into a string given by the path parameter
func ReadFullyFromFile(path string) (string, error) {
	input := ""
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		input = input + line
	}

	return input, nil
}
