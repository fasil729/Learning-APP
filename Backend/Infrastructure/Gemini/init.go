package gemini

import (
	"Brilliant/config"
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	option "google.golang.org/api/option"
)

func GetTextModel() *genai.GenerativeModel {
	// Load environment variables
	env, err := config.Load()
	if err != nil {
		panic(err)
	}
	apiKey := env.Gemin_Api_key
	fmt.Println("API Key: ", apiKey)

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockOnlyHigh,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockMediumAndAbove,
		},

		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
	}
	return model
}

func GetImageModel() *genai.GenerativeModel {
	// Load environment variables
	env, err := config.Load()
	if err != nil {
		panic(err)
	}
	apiKey := env.Gemin_Api_key
	fmt.Println("API Key: ", apiKey)

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Close()

	model := client.GenerativeModel("gemini-pro-vision")

	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockMediumAndAbove,
		},

		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
	}
	return model
}
