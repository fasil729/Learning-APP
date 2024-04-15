package domain

import (
    "github.com/jinzhu/gorm"
)

// UserRole custom type for user roles
type UserRole string

const (
    Teacher UserRole = "teacher"
    Student UserRole = "student"
    Admin   UserRole = "admin"
)

// User model
type User struct {
    gorm.Model
    Username    string      `json:"username" binding:"required"`
    Password    string      `json:"password" binding:"required"`
    Email       string      `json:"email" binding:"required"`
    Role        UserRole    `json:"role" binding:"required" gorm:"type:user_role"`
    RefreshToken  string `json:"refreshToken,omitempty"`
    Topics      []Topic
    Notes       []Note
    Experiments []Experiment
    Quizzes     []Quiz
}

// TableName specifies the table name for the User model
func (User) TableName() string {
    return "users"
}
