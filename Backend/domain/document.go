package domain

import "github.com/jinzhu/gorm"

// Document model
type Document struct {
	gorm.Model
	DocumentName string `json:"document_name" binding:"required"`
	DocumentType string `json:"document_type" binding:"required"`
	FilePath     string `json:"file_path" binding:"required"`
	TopicID      uint   `json:"-"`
}