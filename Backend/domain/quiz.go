package domain

import "github.com/jinzhu/gorm"

// Quiz model
type Quiz struct {
	gorm.Model
	QuizName    string `json:"quiz_name" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID      uint   `json:"-"`
	TopicID     uint   `json:"-"`
}
