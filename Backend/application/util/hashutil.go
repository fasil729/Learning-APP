package util

import (
    "golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}

// CheckPassword checks if the given hashed password matches the given real password.
func CheckPassword(hashedPassword, realPassword string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(realPassword))
    return err == nil
}
