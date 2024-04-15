package dtos

import (
	"Brilliant/domain"
)

// UserDTO represents a user's data in a simplified form.
type UserDTO struct {
    ID            uint  `json:"id"`
    Username      string `json:"username"`
    Email         string `json:"email"`
    Role          domain.UserRole `json:"role"`
    AccessToken   string `json:"accessToken"`
    RefreshToken  string `json:"refreshToken"`
}
