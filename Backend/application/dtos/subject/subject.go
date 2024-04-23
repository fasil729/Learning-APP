package dtos

import "Brilliant/application/dtos/chapter"

// SubjectDTO represents a subject for transport between different parts of the application or between the application and clients.
type SubjectDTO struct {
	ID                  uint              `json:"id"`
	SubjectName         string            `json:"subjectName"`
	UserID              uint              `json:"userId"`
	Chapters            []dtos.ChapterDTO `json:"chapters"`
	TextBookLink        string            `json:"textBookLink"`
	ReferenceBooksLinks []string          `json:"referenceBooksLinks"`
}
