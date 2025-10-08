package main

import (
	"errors"
	"fmt"
	"os"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func readFile(path string) error {
	_, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	return nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	err = readFile("unknown.txt")
	if err != nil {
		fmt.Println("File error:", err)
	}

	fmt.Println("Program finished safely")
}
