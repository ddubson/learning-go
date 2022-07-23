package source

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = cobra.Command{
	Use:   "Learning Go",
	Short: "Learning Go",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
