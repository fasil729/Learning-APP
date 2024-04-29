package gemini

import (
	"context"
	"log"
	"strconv"
	"strings"

	"Brilliant/application/contracts/gemini"
	"Brilliant/application/dtos/exam_prep"

	"github.com/google/generative-ai-go/genai"
)

// GeminiExamPrepHandler implements IExamPrepHandler
type GeminiExamPrepHandler struct {
	model *genai.GenerativeModel
	ctx   context.Context
}

// NewGeminiExamPrepHandler creates a new instance of GeminiExamPrepHandler and returns it as an IExamPrepHandler
func NewGeminiExamPrepHandler() contracts.IExamPrepHandler {
	return &GeminiExamPrepHandler{
		model: GetTextModel(), // Assuming `model` is accessible from this scope
		ctx:   context.Background(),
	}
}

func (eh *GeminiExamPrepHandler) GenerateExamPrep(dto *dtos.GenerateExamPrepDTO) ([]*genai.Candidate, error) {
	// Generate a prompt based on the topics and study time
	topics := strings.Join(dto.Topics, ", ")
	prompt := "Your exam prep topics are: " + topics + ". " +
		"Your prompt is: " + dto.Prompt + ". " +
		"Your study time is: " + strconv.Itoa(dto.ReadTime) + " minutes."
	resp, err := eh.model.GenerateContent(eh.ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return resp.Candidates, nil
}
