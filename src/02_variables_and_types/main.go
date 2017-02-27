package main

import "fmt"

// Global constants
// No `:=` necessary in const block
const (
	constantString = "This is a global constant string\n"
)

var (
	globalString = "This is a global variable string but non-constant"
)

func main() {
	// Long form - variable message of type `string`
	var longMsg string = "This is a long-form string\n"

	// Short form - variable message with dynamic resolution of type
	shortMsg := "This is a short-hand string\n"

	// Int types
	var unsignedInt uint = 50
	luckyNumber := 13

	// Float types
	var pi64 float64 = 3.14
	var pi32 float32 = 3.14
	pi := 3.14                // inferred type
	pi = float64(3.14)        // explicit casting to a type

	// Boolean types
	isTrue := true
	var isNotTrue bool = false

	// Byte types
	b := byte(65)

	// Print outs
	fmt.Print(longMsg)
	fmt.Print(shortMsg)

	fmt.Printf("Lucky number is %d\n", luckyNumber)
	fmt.Printf("Unsigned int is %d\n", unsignedInt)
	fmt.Printf("PI 64 is %.2f\n", pi64)
	fmt.Printf("PI 32 is %.2f\n", pi32)
	fmt.Printf("PI is %.2f\n", pi)
	fmt.Printf("True is %t\n", isTrue)
	fmt.Printf("Is Not True is %t\n", isNotTrue)
	fmt.Printf("Byte 65 in hex is %xh\n", b)

	fmt.Print(constantString)

	fmt.Printf("%s\n", globalString)
	globalString = globalString + ". CHANGED!!"
	fmt.Printf("%s\n", globalString)
}