package domain

import "github.com/jinzhu/gorm"

// StudyRoadmap model
type StudyRoadmap struct {
	gorm.Model
	UserID          uint `json:"-"`
	TopicID         uint `json:"-"`
	Notes           bool `json:"notes" binding:"required"`
	Experiments     bool `json:"experiments" binding:"required"`
	Recommendations bool `json:"recommendations" binding:"required"`
	CreatedMaterial bool `json:"created_material" binding:"required"`
}

