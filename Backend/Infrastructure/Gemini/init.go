package gemini

import (
	"Brilliant/config"
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	option "google.golang.org/api/option"
)


func GetTextModel()  *genai.GenerativeModel {
	// Load environment variables
	env, err := config.Load()
	if err != nil {
		panic(err)
	}
	apiKey := env.Gemin_Api_key

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Close()

	model := client.GenerativeModel("gemini-pro")


	return model
}