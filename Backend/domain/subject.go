package domain

import "github.com/jinzhu/gorm"

type Subject struct {
	gorm.Model
	SubjectName         string
	UserID              uint
	Chapters            []Chapter `gorm:"foreignKey:SubjectID"`
	TextBookLink        string
	ReferenceBookLink string
}
