package domain

import "github.com/jinzhu/gorm"

type Subject struct {
	gorm.Model
	Name               string
	UserID             uint
	Chapters           []Chapter `gorm:"foreignKey:SubjectID"`
	TextBookLink       string
	ReferenceBooksLink string
}
