package contracts

import "Brilliant/application/dtos/exam_prep"

// IExamPrepHandler defines the contract for handling exam preparation in Gemini.
type IExamPrepHandler interface {
	GenerateExamPrep(dto *dtos.GenerateExamPrepDTO) ([]byte, error)
}