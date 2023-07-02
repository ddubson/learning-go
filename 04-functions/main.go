package main

import (
	"fmt"
	"os"
)

func functions() {
	simplePrinter("Hello World!")

	multiVariablePrinter("Hello", "World!", 5)

	if err := returnsErrorPrinter("Hello World!"); err != nil {
		fmt.Print("There was an error\n")
	}

	appendedMsg, err := multiReturnParamPrinter("Hello World")
	if err == nil {
		fmt.Printf("%q\n", appendedMsg)
	}

	deferredExecPrinter("Hello There!")
	deferredExecPrinter("Another message")

	filePrinter("Write this message to a file")

	namedReturnValPrinter("Yet Another Message")

	listParamPrinter("Can", "You", "Hear", "Me", "Now!")
}

func simplePrinter(msg string) {
	fmt.Printf("%s\n", msg)
}

// - Condensed param declaration
func multiVariablePrinter(msg, msg2 string, repeat int) {
	for repeat > 0 {
		fmt.Printf("%s + %s\n", msg, msg2)
		repeat--
	}
}

func returnsErrorPrinter(msg string) error {
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func multiReturnParamPrinter(msg string) (string, error) {
	msg += "\n"
	_, err := fmt.Printf(msg)
	return msg, err
}

func deferredExecPrinter(msg string) error {
	// defer execution until end of function
	defer fmt.Print("exiting printer5\n")

	_, err := fmt.Printf("%s\n", msg)
	return err
}

func filePrinter(msg string) error {
	f, err := os.Create("static/05_functions.example")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(msg)
	return err
}

func namedReturnValPrinter(msg string) (e error) {
	_, e = fmt.Printf("%s\n", msg)
	return e
}

func listParamPrinter(msgs ...string) {
	for _, msg := range msgs {
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	functions()
}
