package main

import "fmt"

func main() {
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
