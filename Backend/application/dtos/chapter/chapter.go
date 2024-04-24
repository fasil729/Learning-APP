package dtos

import "Brilliant/application/dtos/lesson"

type ChapterDTO struct {
	SubjectID   uint             `json:"id"`
	ChapterName string           `json:"chapterName"`
	Lessons     []dtos.LessonDTO `json:"lessons"`
}
