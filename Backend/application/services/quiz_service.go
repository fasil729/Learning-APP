package services

import (
	"errors"
	"math/rand"
	"time"
	"Brilliant/application/dtos/quiz"
)

type QuizService struct{}





func NewQuizService() *QuizService {
	return &QuizService{}
}

func (qs *QuizService) GenerateQuiz(dto *dtos.GenerateQuizDTO) ([]dtos.Quiz, error) {
	if len(dto.Topics) == 0 {
		return nil, errors.New("topics list is empty")
	}

	rand.Seed(time.Now().UnixNano())

	quizzes := make([]dtos.Quiz, dto.NumberOfQuizzes)
	for i := 0; i < dto.NumberOfQuizzes; i++ {
		// For simplicity, we'll generate sample quizzes here
		quizzes[i] = dtos.Quiz{
			Question:   "Sample question about the circulatory system",
			Options:    []dtos.Option{{Text: "Option 1", IsAnswer: false}, {Text: "Option 2", IsAnswer: false}, {Text: "Option 3", IsAnswer: true}, {Text: "Option 4", IsAnswer: false}},
			Explanation: "Explanation for the correct answer",
		}
	}

	return quizzes, nil
}

