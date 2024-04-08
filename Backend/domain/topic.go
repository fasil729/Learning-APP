package domain

import "github.com/jinzhu/gorm"

// Topic model
type Topic struct {
	gorm.Model
	TopicName    string    `json:"topic_name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	UserID       uint      `json:"-"`
	Documents    []Document
	Notes        []Note
	Experiments  []Experiment
	Quizzes      []Quiz
}
