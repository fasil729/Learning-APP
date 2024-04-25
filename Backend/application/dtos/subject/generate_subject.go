package dtos

type GenerateSubjectDTO struct {
	SubjectID   uint   `json:"subjectId" binding:"required"`
	ChapterID   uint   `json:"chapterId" binding:"required"`
	SubjectName string `json:"subjectName" binding:"required"`
	UserID      uint   `json:"userId" binding:"required"`
}
