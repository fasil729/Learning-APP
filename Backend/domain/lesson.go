package domain

import "github.com/jinzhu/gorm"

type Lesson struct {
	gorm.Model
	Name        string
	ChapterID   uint
	ContentLink string
}
