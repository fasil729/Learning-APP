package contracts

import (
	"Brilliant/application/dtos/exam_prep"
	// "github.com/google/generative-ai-go/genai"
)

// IExamPrepHandler defines the contract for handling exam preparation in Gemini.
type IExamPrepHandler interface {
	GenerateExamPrep(dto *dtos.GenerateExamPrepDTO) ([]byte, error)
}
