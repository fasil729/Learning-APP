package domain

import "github.com/jinzhu/gorm"

type Lesson struct {
	gorm.Model
	LessonName  string `json:"name" binding:"required"`
	ChapterID   uint   `json:"chapter_id"`
	ContentLink string `json:"content_link"`
}
