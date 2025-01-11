package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func embeddingAPI() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get Hugging Face API Key from environment variable
	apiKey := os.Getenv("HF_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: HUGGING_FACE_API_KEY environment variable is not set")
	}

	// Hugging Face model endpoint (e.g., "https://api-inference.huggingface.co/models/sentence-transformers/all-MiniLM-L6-v2")
	apiURL := "https://api-inference.huggingface.co/models/mixedbread-ai/mxbai-embed-large-v1"

	// Text input for embedding
	text := "This is an example sentence to embed."

	// Create the request payload
	payload := map[string]interface{}{
		"inputs": "Today is a sunny day and I will get some ice cream.",
	}

	// Serialize payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Failed to serialize payload: %v", err)
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Add headers
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Check for errors in response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: %s\nResponse: %s", resp.Status, body)
	}

	// Deserialize the response
	var embeddingResponse []float64
	err = json.Unmarshal(body, &embeddingResponse)
	if err != nil {
		log.Fatalf("Failed to parse response JSON: %v", err)
	}

	// Print the embeddings
	fmt.Printf("Embeddings for input text: %s\n", text)
	fmt.Println(embeddingResponse)
}

func openAIChat() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("Error: OPENAI_API_KEY environment variable is not set")
	}

	// Create a new OpenAI client
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	// Create a chat completion request
	chatCompletion, err := client.Chat.Completions.New(context.TODO(), openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Say this is a test"),
		}),
		Model: openai.F(openai.ChatModelGPT4oMini), // Correct model name
	})
	if err != nil {
		log.Fatalf("Error creating chat completion: %v", err)
	}

	// Print the response
	fmt.Println("Response:", chatCompletion.Choices[0].Message.Content)
}

func main() {
	embeddingAPI()
	openAIChat()
}
