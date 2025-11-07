/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

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
	base_url := "https://huggingface.co/BAAI/bge-small-en-v1.5/resolve/main/onnx/model.onnx"

	resp, err := http.Get(base_url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	f, err := os.Create("/tmp/model.onnx")
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = io.Copy(f, resp.Body)

	return err
}
