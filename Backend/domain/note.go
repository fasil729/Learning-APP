package domain

import "github.com/jinzhu/gorm"

// Note model
type Note struct {
	gorm.Model
	NoteContent string `json:"note_content" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
	UserID      uint   `json:"-"`
	TopicID     uint   `json:"-"`
}
