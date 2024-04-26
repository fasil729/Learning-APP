package gemini

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"Brilliant/application/contracts/gemini"
)

// GeminiNoteHandler implements INoteHandler
type GeminiNoteHandler struct {
	model *genai.GenerativeModel
	ctx   context.Context
}

// NewGeminiNoteHandler creates a new instance of GeminiNoteHandler and returns it as an INoteHandler
func NewGeminiNoteHandler() contracts.INoteHandler {
	return &GeminiNoteHandler{
		model: GetTextModel(), // Assuming `model` is accessible from this scope
		ctx:   context.Background(),
	}
}
func (nh *GeminiNoteHandler) AddNoteForChapter(noteContent string, chapterName string) error {
	// Generate a prompt to add a note for the chapter
	prompt := "Add a note for the chapter " + chapterName + ". " + "Note: " + noteContent
	_, err := nh.model.GenerateContent(nh.ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (nh *GeminiNoteHandler) AddNoteForChapterFromImage(imageData []byte, chapterName string) error {
	// Generate a prompt to add a note for the chapter from an image
	prompt := "Add a note for the chapter " + chapterName + " from the provided image."
	_, err := nh.model.GenerateContent(nh.ctx, genai.ImageData("jpeg", imageData), genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}