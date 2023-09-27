package main

import (
	"fmt"
	"os"
)

// Global constants
// No `:=` necessary in const block
const (
	globalImmutableString = "This is a global constant string\n"
)

var (
	globalMutableString = "This is a global variable string but non-constant"
)

func basicDataTypes() {
	// Long form variable declaration
	// Example of a "string literal"
	var longMsg string = "This is a long-form string\n"

	// Short form - variable message with type inference
	// **Implicit initialization syntax ":="**
	shortMsg := "This is a short-hand string\n"

	// Raw string literal example
	rawStringLiteral := `hi 
		there! This allows for multi-line strings`
	fmt.Printf(rawStringLiteral)

	// Integer types
	var unsignedInt uint = 50
	luckyNumber := 13

	// Float types
	var pi64 float64 = 3.14
	var pi32 float32 = 3.14
	pi := 3.14         // inferred type
	pi = float64(3.14) // explicit casting to a type

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

	fmt.Print(globalImmutableString)

	fmt.Printf("%s\n", globalMutableString)
	globalMutableString = globalMutableString + ". CHANGED!!"
	fmt.Printf("%s\n", globalMutableString)
}

func mutableVars() {
	// Variables are mutable.
	var i int32
	i = 1
	fmt.Println(i)
	i = 2
	fmt.Println(i)
	i = 3
	fmt.Println(i)

	j := 1
	fmt.Println(j)
	j = 2
	fmt.Println(j)
	j = 3
	fmt.Println(j)
}

func pointerDataTypes() {
	fmt.Println("! Pointers !")
	// Declare a pointer data type - it has to be initialized in order for contents to be set
	var firstName *string = new(string)

	// De-reference the pointer, and set the variable
	*firstName = "Arthur"

	fmt.Printf("The memory address of firstName: %xh\n", firstName)
	fmt.Printf("The contents of firstName: %s\n", *firstName)

	// Short-hand
	lastName := "Williams"
	// Using the "address-of" & operator
	pointerToLastName := &lastName
	fmt.Printf("Last name is %s and pointer to it is at %xh\n", lastName, pointerToLastName)
}

func byteSlices() {
	f, err := os.Open("static/02_variables_and_types.txt")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	b := make([]byte, 100)
	n, err := f.Read(b)
	fmt.Printf("%d: % x\n", n, b)

	stringVersion := string(b)
	fmt.Printf("%s\n", stringVersion)
}

func maps() {
	fmt.Println("\n------- MAPS ----------")

	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

	fmt.Printf("Days in February: %d\n", dayMonths["Feb"])
	days, ok := dayMonths["January"]
	if !ok {
		fmt.Print("Can't get days for January\n")
	} else {
		fmt.Printf("%d\n", days)
	}

	// Iterating over a map
	for month, days := range dayMonths {
		fmt.Printf("%s has %d days\n", month, days)
	}

	// Delete an item from a map
	delete(dayMonths, "Feb")
}

func arrays() {
	fmt.Println("\n------- ARRAYS ----------")

	// Simple array with unknown # of string values
	words := [...]string{"the", "quick", "brown", "fox"}
	fmt.Printf("%s\n", words[2])
	// Arrays are passed in via value -> whole array is copied to the stackframe of callee

	// Slices
	wordsSlice := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	// Slices are passed by reference
	printer(wordsSlice)
	printer(wordsSlice[5:])

	wordsDynSlic := make([]string, 4)
	fmt.Printf("len: %d, cap: %d\n", len(wordsDynSlic), cap(wordsDynSlic))
	wordsDynSlic[0] = "Do"
	wordsDynSlic[1] = "I"
	wordsDynSlic[2] = "Amuse"
	wordsDynSlic[3] = "You"
	printer(wordsDynSlic)
}

func stringFunctions() {
	fmt.Println("\n------- STRING MANIPULATION ----------")
	atoz := "the quick brown fox jumps over the lazy dog\n"

	// Substring - get first 9 characters of string
	fmt.Printf("%s\n", atoz[0:9])
	fmt.Printf("%s\n", atoz[15:19])
	fmt.Printf("%s\n", atoz[15:])

	for _, r := range atoz {
		fmt.Printf("%c ", r)
	}

	literalString := `Everything here is \t interpreted \n literally.`
	fmt.Printf("\n%s\n", literalString)
}

func printer(w []string) {
	for _, word := range w {
		fmt.Printf("%s ", word)
	}
	fmt.Println()
}

func main() {
	basicDataTypes()
	mutableVars()
	pointerDataTypes()
	stringFunctions()
	arrays()
	maps()
	byteSlices()
	TypeConversionShowcase()
}
