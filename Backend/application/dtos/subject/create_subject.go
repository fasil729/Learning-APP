package dtos

// SubjectDTO represents a subject for transport between different parts of the application or between the application and clients.
type CreateSubjectDTO struct {
	SubjectName string `json:"subjectName"`
}
