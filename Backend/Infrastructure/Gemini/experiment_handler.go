package gemini

import (
	"context"
	"log"

	"github.com/google/generative-ai-go/genai"
	"Brilliant/application/contracts/gemini"
)

// GeminiExperimentHandler implements IExperimentHandler
type GeminiExperimentHandler struct {
	model *genai.GenerativeModel
	ctx   context.Context
}

// NewGeminiExperimentHandler creates a new instance of GeminiExperimentHandler and returns it as an IExperimentHandler
func NewGeminiExperimentHandler() contracts.IExperimentHandler {
	return &GeminiExperimentHandler{
		model: model, // Assuming `model` is accessible from this scope
		ctx:   context.Background(),
	}
}

func (eh *GeminiExperimentHandler) GetExperimentsForChapter(chapterName string) ([]*genai.Candidate, error) {
	resp, err := eh.model.GenerateContent(eh.ctx, genai.Text("List possible experiments for the chapter "+chapterName))
	if err != nil {
		log.Fatal(err)
		return resp.Candidates, err
	}

	return resp.Candidates, nil
}

func (eh *GeminiExperimentHandler) GetExperimentContent(experimentName string) ([]*genai.Candidate, error) {
	resp, err := eh.model.GenerateContent(eh.ctx, genai.Text("Generate content for the experiment "+experimentName))
	if err != nil {
		log.Fatal(err)
		return resp.Candidates, err
	}

	return resp.Candidates, nil
}
