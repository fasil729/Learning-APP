package domain

import "github.com/jinzhu/gorm"

// User model
type User struct {
    gorm.Model
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    Email    string `json:"email" binding:"required"`
    Role     string `json:"role" binding:"required" gorm:"type:ENUM('teacher', 'student', 'admin')"`
    Topics   []Topic
    Notes    []Note
    Experiments []Experiment
    Quizzes  []Quiz
}
