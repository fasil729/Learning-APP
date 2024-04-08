package domain

import "github.com/jinzhu/gorm"

// Experiment model
type Experiment struct {
	gorm.Model
	ExperimentName string `json:"experiment_name" binding:"required"`
	Description    string `json:"description" binding:"required"`
	UserID         uint   `json:"-"`
	TopicID        uint   `json:"-"`
}
