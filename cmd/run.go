/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run a configured server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Running '%v'", args[0])
		execCmd := exec.Command("ollama", "run", args[0])
		value, err := execCmd.CombinedOutput()
		if err != nil {
			exitErr, ok := err.(*exec.ExitError)
			if ok {
				fmt.Printf("ollama run fail. Did you run serve? err: %s\n", exitErr.Stderr)
			} else {
				fmt.Println("ollama run fail.")
			}
		}
		fmt.Println("Ollama run success.", value)
	},
}

func init() {
	serverCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().String("model", "", "Name of the model")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
