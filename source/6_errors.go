package source

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	errorEmptyString = errors.New("Don't want to print empty string")
)

func errorsExample() {
	if err := printMsg(""); err != nil {
		fmt.Printf("printer failed: %s\n", err)
		os.Exit(1)
	}
}

func printMsg(msg string) error {
	if msg == "" {
		return errorEmptyString
	}

	_, err := fmt.Printf("%s\n", msg)
	return err
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "err",
		Short: "Errors demonstration module",
		Run: func(cmd *cobra.Command, args []string) {
			errorsExample()
		},
	})
}
