package dtos

import "Brilliant/application/dtos/chapter"

// SubjectDTO represents a subject for transport between different parts of the application or between the application and clients.
type SubjectDTO struct {
	ID                 uint              `json:"id"`
	UserID             uint              `json:"userId"` // The ID of the user who created the subject
	SubjectName        string            `json:"subjectName"`
	Chapters           []dtos.ChapterDTO `json:"chapters"`
	TextBookLink       string            `json:"textBookLink"`
	ReferenceBooksLink string            `json:"referenceBooksLink"`
}
