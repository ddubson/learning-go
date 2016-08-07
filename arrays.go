package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		arraySize, _ := strconv.Atoi(input.Text())
		array1 := make([]string, arraySize)
		input.Scan()

		array1 = strings.Split(input.Text(), " ")
		for x := arraySize - 1; x >= 0; x-- {
			fmt.Printf("%s ", array1[x])
		}
	}
}
