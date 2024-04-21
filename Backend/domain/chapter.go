package domain

import "github.com/jinzhu/gorm"

type Chapter struct {
	gorm.Model
	Name        string
	SubjectID   uint
	Note        Note         `gorm:"foreignKey:ChapterID"`
	Lessons     []Lesson     `gorm:"foreignKey:ChapterID"`
	Experiments []Experiment `gorm:"foreignKey:ChapterID"`
}
