package dtos

type CreateChapterDTO struct {
	SubjectID   uint   `json:"id"`
	ChapterName string `json:"chapterName"`
}
