package services

import (
	gemini "Brilliant/Infrastructure/Gemini"
	"Brilliant/application/contracts/gemini"
	"Brilliant/application/dtos/quiz"
	"errors"
)

type QuizService struct{
	geminiQuizHandler contracts.IQuizHandler
}





func NewQuizService() *QuizService {
	return &QuizService{
		geminiQuizHandler: gemini.NewGeminiQuizHandler(),
	}
}

func (qs *QuizService) GenerateQuiz(dto *dtos.GenerateQuizDTO) ([]dtos.Quiz, error) {
	if len(dto.Topics) == 0 {
		return nil, errors.New("topics list is empty")
	}


	quizzes, err := qs.geminiQuizHandler.GenerateQuiz(dto)
	if err != nil {
		return nil, err
	}

	return quizzes, nil
}

