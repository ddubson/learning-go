package main

import (
	"fmt"
)

type Line string
type Stanza []Line
type Poem []Stanza

func main() {
	fmtPackage()
}

func fmtPackage() {
	p := Poem{{"Grave men, near death, who see with blinding sight", "Blind eyes could blaze like meteors and be gay", "Rage, rage against the dying of the light"}, {"And you, my father, there on the sad height,", "Curse, bless, me now with your fierce tears, I pray.", "Do not go gentle into that good night.", "Rage, rage against the dying of the light."}}

	fmt.Printf("%v\n\n", p)  // Print out the whole value for a complex type
	fmt.Printf("%#v\n\n", p) // Prints out the whole value + types within
	fmt.Printf("%T\n\n", p)    // Prints out the type of value
	fmt.Printf("%q\n\n", p[0][0])  // Print in quotes the line
}