package domain

import "github.com/jinzhu/gorm"

// Experiment model
type Experiment struct {
	gorm.Model
	ExperimentName string `json:"experiment_name" binding:"required"`
	SubjectID        uint   `json:"topic_id"`
	ChapterID      uint   `json:"chapter_id"`
	ContentLink    string `json:"content_link"`
}