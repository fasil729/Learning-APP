package dtos

type UpdateUserDTO struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    // Add other fields that can be updated
}