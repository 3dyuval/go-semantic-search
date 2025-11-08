package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
)

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
