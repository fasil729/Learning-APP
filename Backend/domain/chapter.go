package domain

import "github.com/jinzhu/gorm"

type Chapter struct {
	gorm.Model
	Name        string       `json:"name" binding:"required"`
	SubjectID   uint         `json:"subject_id"`
	Note        Note         `gorm:"foreignKey:ChapterID"`
	Lessons     []Lesson     `gorm:"foreignKey:ChapterID"`
	Experiments []Experiment `gorm:"foreignKey:ChapterID"`
	NoteLink    string
}
