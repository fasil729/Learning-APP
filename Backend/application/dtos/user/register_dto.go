package dtos

// RegisterUserDTO represents the data required to register a new user.
type RegisterDTO struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    Email    string `json:"email" binding:"required"`
    Role     string `json:"role" binding:"required" enum:"student teacher,admin"`
}