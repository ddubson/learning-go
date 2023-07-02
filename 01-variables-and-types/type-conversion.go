package main

import "fmt"

// TypeConversionShowcase
// see https://go.dev/ref/spec#Conversions
func TypeConversionShowcase() {
	fmt.Println("Type conversion showcase")

	fmt.Println(float32(2.718281828)) // 2.718281828 of type float32
	fmt.Println(complex128(1))        // 1.0 + 0.0i of type complex128
	fmt.Println(float32(0.49999999))  // 0.5 of type float32
	fmt.Println(float64(-1e-1000))    // 0.0 of type float64
	fmt.Println(string('x'))          // "x" of type string
	fmt.Println(string(0x266c))       // "♬" of type string
	fmt.Println(string([]byte{'a'}))  // not a constant: []byte{'a'} is not a constant
	fmt.Println((*int)(nil))          // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
	// fmt.Println(myString("foo" + "bar")) // "foobar" of type myString
	// fmt.Println(uint(iota))	      // iota value of type uint
	// fmt.Println(int(1.2))         // illegal: 1.2 cannot be represented as an int
	// fmt.Println(string(65.0))     // illegal: 65.0 is not an integer constant
}
