package main

import (
	"fmt"
	"os"
)

func simple_printer(msg string) {
	fmt.Printf("%s\n", msg)
}

// - Condensed param declaration
func multi_variable_printer(msg, msg2 string, repeat int) {
	for repeat > 0 {
		fmt.Printf("%s + %s\n", msg, msg2)
		repeat--
	}
}

func returns_error_printer(msg string) error {
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func multi_return_param_printer(msg string) (string, error) {
	msg += "\n"
	_, err := fmt.Printf(msg)
	return msg, err
}

func deferred_exec_printer(msg string) error {
	// defer execution until end of function
	defer fmt.Print("exiting printer5\n")

	_, err := fmt.Printf("%s\n", msg)
	return err
}

func file_printer(msg string) error {
	f, err := os.Create("static/05_functions.example")
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(msg)
	return err
}

func named_return_val_printer(msg string) (e error) {
	_, e = fmt.Printf("%s\n", msg)
	return e
}

func list_param_printer(msgs ...string) {
	for _, msg := range msgs {
		fmt.Printf("%s\n", msg)
	}
}

func main() {
	simple_printer("Hello World!")

	multi_variable_printer("Hello", "World!", 5)

	if err := returns_error_printer("Hello World!"); err != nil {
		fmt.Print("There was an error\n")
	}

	appendedMsg, err := multi_return_param_printer("Hello World")
	if err == nil {
		fmt.Printf("%q\n", appendedMsg)
	}

	deferred_exec_printer("Hello There!")
	deferred_exec_printer("Another message")

	file_printer("Write this message to a file")

	named_return_val_printer("Yet Another Message")

	list_param_printer("Can", "You", "Hear", "Me", "Now!")
}
