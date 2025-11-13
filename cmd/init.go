/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Download model for vectorization",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Downloadind model")
		return getModel()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func getModel() error {
	cmd := exec.Command("ollama", "list")
	value := cmd.Run()

	if value != nil {
		fmt.Println("Found ollama models: %", value)
		return value
	}

	fmt.Println("No ollama models found")

	return value
}
