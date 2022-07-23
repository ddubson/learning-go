package source

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func conditionals() {
	if true {
		fmt.Print("Something here\n")
	}

	// Shorthand for checking if error exists during printing.
	if numberOfBytes, err := fmt.Print("Hello\n"); err != nil {
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d characters\n", numberOfBytes)
	}

	// Switch statement
	// - No break statements!
	// You can use `fallthrough` to continue down the chain
	n, err := fmt.Print("Print again\n")
	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Print("No bytes output\n")
	case n != 12:
		fmt.Printf("Wrong number of chars %d\n", n)
	default:
		fmt.Print("OK!\n")
	}

	switchExampleVowelCount()
}

func switchExampleVowelCount() {
	atoz := "the quick brown fox jumps over the lazy dog"
	vowels := 0
	consonants := 0
	zeds := 0

	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels += 1
		case 'z':
			zeds += 1
			fallthrough
		default:
			consonants += 1
		}
	}

	fmt.Printf("Vowels: %d, Consonants: %d (with %d zeds)\n", vowels, consonants, zeds)
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "conditionals",
		Short: "Conditionals demonstration module",
		Run: func(cmd *cobra.Command, args []string) {
			conditionals()
		},
	})
}
