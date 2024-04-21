// ExtractUserDetails extracts the user details from the token.
package util

import (
	dtos "Brilliant/application/dtos/user"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// CreateAccessToken creates a new access token for the user.
func CreateAccessToken(userID uint, username, email, role, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := jwt.MapClaims{
		"sub":      userID,
		"username": username,
		"email":    email,
		"role":     role,
		"exp":      exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, nil
}

// CreateRefreshToken creates a new refresh token for the user.
func CreateRefreshToken(userID uint, username, email, role, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := jwt.MapClaims{
		"sub":      userID,
		"username": username,
		"email":    email,
		"role":     role,
		"exp":      exp,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return rt, nil
}

// IsTokenValid checks if a token is valid.
func IsTokenValid(tokenString string, secret string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return false
	}
	return true
}


// ExtractUserID extracts the user ID from the token.
func ExtractUserID(tokenString string, secret string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}
	sub, ok := claims["sub"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid token sub claim")
	}
	return uint(sub), nil
}

// ExtractUserDetails extracts the user details from the token.
func ExtractUserDetails(tokenString string, secret string) (*dtos.AuthDetailDTO, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	sub, ok := claims["sub"].(float64)
	if !ok {
		return nil, fmt.Errorf("invalid token sub claim")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token username claim")
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token email claim")
	}
	role, ok := claims["role"].(string)
	if !ok {
		return nil, fmt.Errorf("invalid token role claim")
	}
	return &dtos.AuthDetailDTO{
		UserID:   uint(sub),
		Username: username,
		Email:    email,
		Role:     role,
	}, nil
}
