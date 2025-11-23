package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var embedCmd = &cobra.Command{
	Use:   "embed",
	Short: "Create embedding request",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("embed called")
		value, err := getEmbedding(args[0])
		if err != nil {
			println("Error", err)
		}

		println("ok", value)
	},
}

func init() {
	rootCmd.AddCommand(embedCmd)

	serverCmd.PersistentFlags().String("input", "", "The content to embed")
}

type EmbedRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}
type EmbedResponse struct {
	Embedding []float64 `json:"embedding"`
}

func getEmbedding(text string) ([]float64, error) {
	reqBody := EmbedRequest{
		Model:  "nomic-embed-text",
		Prompt: text,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post("http://localhost:11434/api/embeddings",
		"application/json",
		bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result EmbedResponse

	json.NewDecoder(resp.Body).Decode(&result)

	return result.Embedding, nil
}
