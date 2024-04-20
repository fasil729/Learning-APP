package domain

import "github.com/jinzhu/gorm"

// Document model
type Lesson struct {
    gorm.Model
    Name      string    `json:"name" binding:"required"`
    ChapterID uint  `json:"chapter_id"`
	contentLink   string `json:"content_link"`
}