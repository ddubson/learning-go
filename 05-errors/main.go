package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	errorEmptyString = errors.New("Don't want to print empty string")
)

func errorsExample() {
	if err := printMsg(""); err != nil {
		fmt.Printf("printer failed: %s\n", err)
		os.Exit(1)
	}
}

func printMsg(msg string) error {
	if msg == "" {
		return errorEmptyString
	}

	_, err := fmt.Printf("%s\n", msg)
	return err
}

func main() {
	errorsExample()
}
