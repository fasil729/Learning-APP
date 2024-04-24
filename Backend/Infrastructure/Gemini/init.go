package gemini

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
);

var model *genai.GenerativeModel

func init() {
	
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
	log.Fatal(err)
	}
	defer client.Close()

	model = client.GenerativeModel("Brilliant_Model")

}