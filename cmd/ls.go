/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List moders",
	Run: func(cmd *cobra.Command, args []string) {
		cmdExec := exec.Command("ollama", "list")
		value, err := cmdExec.CombinedOutput()

		errOut, ok := err.(*exec.ExitError)
		if ok {
			fmt.Println("Error retrieving ollama models: ", string(value), errOut.String())
			return
		}

		if strings.HasPrefix("Error: ", string(value)) {
			fmt.Println("No ollama models found")
			return
		}

		fmt.Printf("Found ollama models: %s", value)
	},
}

func init() {
	serverCmd.AddCommand(lsCmd)
}
