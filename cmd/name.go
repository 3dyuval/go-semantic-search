package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Hello world with user name",
	Run: func(cmd *cobra.Command, args []string) {
		run(args)
	},
}

func run(args []string) {
	fmt.Println("Hello " + args[0] + "!")
}

func init() {
	rootCmd.AddCommand(nameCmd)

	// TODO investigate PersistentFlags vs local flags
	nameCmd.PersistentFlags().String("name", "world", "Your name")
	nameCmd.LocalFlags().StringP("local-name", "n", "world", "Your name")
}
