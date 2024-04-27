package dtos

import ( 
	"Brilliant/domain"
	"Brilliant/application/contracts/gemini"
)

// SubjectWithChaptersDTO represents a subject along with its chapters.
type SubjectWithChaptersDTO struct {
	Subject  *domain.Subject
	Chapters []contracts.Chapter
}