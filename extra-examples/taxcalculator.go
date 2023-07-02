package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	// declare variables, initialized to default
	var mealCost, tipPercent, taxPercent float64

	if input.Scan() {
		// convert string to float 64-bit
		mealCost, _ = strconv.ParseFloat(input.Text(), 64)
	}

	if input.Scan() {
		tipPercent, _ = strconv.ParseFloat(input.Text(), 64)
	}

	if input.Scan() {
		taxPercent, _ = strconv.ParseFloat(input.Text(), 64)
	}

	tip := mealCost * (tipPercent / 100)
	tax := mealCost * (taxPercent / 100)
	// Round function does not come built-in Go, so we do int(f+0.5)
	fmt.Printf("The total meal cost is %d dollars.\n", int(mealCost + tip + tax + 0.5))
}