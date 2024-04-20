package domain

import "github.com/jinzhu/gorm"

// Note model
type Note struct {
	gorm.Model
	ChapterID   uint
	ContentLink string
}
