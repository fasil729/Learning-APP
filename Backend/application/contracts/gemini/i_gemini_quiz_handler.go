package contracts

import "Brilliant/application/dtos/quiz"

// IQuizHandler defines the contract for handling quizzes in Gemini.
type IQuizHandler interface {
	GenerateQuiz(dto *dtos.GenerateQuizDTO) ([]dtos.Quiz, error)
}