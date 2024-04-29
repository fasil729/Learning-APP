package dtos

type GenerateQuizDTO struct {
	Topics       []string `json:"topics"`
	Prompt       string   `json:"prompt"`
	NumberOfQuizzes int    `json:"number_of_quizzes"`
}