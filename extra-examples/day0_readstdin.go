package main

import (
	"fmt"
	"os"
	"bufio"
)

/**
 HackerRank 30 Days of Code : Day 0

 Read from command line and print it out.
 */
func main() {
	fmt.Printf("Hello World\n")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}