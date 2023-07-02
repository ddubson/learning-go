package main

import (
	"bufio"
	"fmt"
	"os"
)

func SimpleStdinRead() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	fmt.Println(text)
}

func main() {
	SimpleStdinRead()
}
