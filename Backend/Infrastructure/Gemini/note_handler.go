package gemini

import (
	"context"
	"encoding/json"
	"log"

	contracts "Brilliant/application/contracts/gemini"

	"github.com/google/generative-ai-go/genai"
)

// GeminiNoteHandler implements INoteHandler
type GeminiNoteHandler struct {
	model      *genai.GenerativeModel
	ImageModel *genai.GenerativeModel
	ctx        context.Context
}

// NewGeminiNoteHandler creates a new instance of GeminiNoteHandler and returns it as an INoteHandler
func NewGeminiNoteHandler() contracts.INoteHandler {
	return &GeminiNoteHandler{
		model:      GetTextModel(),
		ImageModel: GetImageModel(), // Assuming `model` is accessible from this scope
		ctx:        context.Background(),
	}
}
func (nh *GeminiNoteHandler) AddNoteForChapter(previousContent string, noteContent string, chapterName string) ([]byte, error) {
	// Generate a prompt to add a note for the chapter
	prompt := "merge a new note and  the previous note which are both about the chapter " + chapterName + ". " + "PreviousNote: " + previousContent + ". " + "NewNote: " + noteContent + "." + "The merged note should be coherent and informative. and Give the merged markdown note without replacing the previous content"
	response, err := nh.model.GenerateContent(nh.ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseContent := response.Candidates[0].Content.Parts[0]

	// Convert the response content to byte
	responseContentJSON, _ := json.Marshal(responseContent)

	// Convert the response content from byte to string
	var responseContentString string
	err = json.Unmarshal(responseContentJSON, &responseContentString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return []byte(responseContentString), nil

}

func (nh *GeminiNoteHandler) AddNoteForChapterFromImage(previousContent string, imageData []byte, chapterName string) ([]byte, error) {
	// Generate a prompt to add a note for the chapter from an image
	prompt := "merge the text from the provided image with the previous content which are both the chapter " + chapterName + "previousContent:" + previousContent + "The note should be coherent and informative. and Give the merged markdown note without replacing the previous content"
	response, err := nh.ImageModel.GenerateContent(nh.ctx, genai.ImageData("png", imageData), genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	responseContent := response.Candidates[0].Content.Parts[0]

	// Convert the response content to byte
	responseContentJSON, _ := json.Marshal(responseContent)

	// Convert the response content from byte to string
	var responseContentString string
	err = json.Unmarshal(responseContentJSON, &responseContentString)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return []byte(responseContentString), nil
}
