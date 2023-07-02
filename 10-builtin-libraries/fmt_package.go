package main

import "fmt"

func FmtPackageShowcase() {
	fmt.Printf("%v\n\n", poem)       // Print out the whole value for a complex type
	fmt.Printf("%#v\n\n", poem)      // Prints out the whole value + types within
	fmt.Printf("%T\n\n", poem)       // Prints out the type of value
	fmt.Printf("%q\n\n", poem[0][0]) // Print in quotes the line
}
