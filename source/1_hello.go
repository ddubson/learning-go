package source

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Hello returns a greeting as a string
func Hello() string {
	return "Hello, World!"
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "hello",
		Short: "Welcome!",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print(Hello())
		},
	})
}
