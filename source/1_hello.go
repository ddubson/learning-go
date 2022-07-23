package source

import (
	"fmt"
	"github.com/spf13/cobra"
)

func hello() {
	fmt.Print("hello World!")
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "hello",
		Short: "Welcome!",
		Run: func(cmd *cobra.Command, args []string) {
			hello()
		},
	})
}
