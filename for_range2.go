package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		digit, _ := strconv.Atoi(input.Text())
		for i := 1; i <= 10; i++ {
			fmt.Printf("%d x %d = %d\n", digit , i, digit * i)
		}
	}
}