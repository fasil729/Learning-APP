package domain

import "github.com/jinzhu/gorm"

// Topic model
type Subject struct {
	gorm.Model
	Name                string `json:"name" binding:"required"`
	UserID              uint  `json:"user_id"`
	Chapters            []Chapter `gorm:"foreignKey:SubjectID"`
	TextBookLink        string `json:"text_book_link"`
	ReferenceBooksLink  string `json:"reference_books_links"`	
  }
