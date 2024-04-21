package domain

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
  
	ChapterID uint
	contentLink string
}
