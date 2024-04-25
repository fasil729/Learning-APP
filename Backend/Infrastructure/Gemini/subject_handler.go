package gemini

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"Brilliant/application/contracts/gemini"
)

// GeminisubjectHandler implements ISubjectHandler
type GeminiSubjectHandler struct {
	model *genai.GenerativeModel
	ctx  context.Context
}

// NewGeminiSubjectHandler creates a new instance of GeminisubjectHandler and returns it as an IGeminiSubjectHandler
func NewGeminiSubjectHandler() contracts.IGeminiSubjectHandler {
	return &GeminiSubjectHandler{
		model: model, // Assuming `model` is accessible from this scope
		ctx:  context.Background(),
	}
}

func (sh *GeminiSubjectHandler) GenerateTableOfContent(subjectName string) ([]*genai.Candidate, error) {
	resp, err := sh.model.GenerateContent(sh.ctx, genai.Text("Generate a table of content for the subject "+subjectName))
	if err != nil {
		log.Fatal(err)
		return resp.Candidates, err
	}

	return resp.Candidates, nil
}

func (sh *GeminiSubjectHandler) GenerateLessonDetailContent(lessonName string) ([]*genai.Candidate, error) {
	resp, err := sh.model.GenerateContent(sh.ctx, genai.Text("Generate detailed content for the lesson "+lessonName))
	if err != nil {
		log.Fatal(err)
		return resp.Candidates, err
	}

	return resp.Candidates, nil
}
