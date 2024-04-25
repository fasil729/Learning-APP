package contracts

import "github.com/google/generative-ai-go/genai"

// Define the ISubjectHandler interface in the contracts package
type IGeminiSubjectHandler interface {
	GenerateTableOfContent(subjectName string) ([]*genai.Candidate, error)
	GenerateLessonDetailContent(lessonName string) ([]*genai.Candidate, error)
}