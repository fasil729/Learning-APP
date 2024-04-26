package gemini

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	option "google.golang.org/api/option"
)


func GetTextModel()  *genai.GenerativeModel {
	apiKey := "AIzaSyBRgGXo2clQ41IpQsjzWET2xYQyBcmtjW4"
	if apiKey == "" {
		log.Fatal("Please set the GEMINI_API_KEY environment variable")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Close()

	model := client.GenerativeModel("gemini-pro")


	return model
}
