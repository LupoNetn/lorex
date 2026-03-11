package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// FormatValidationError handles translating raw validator errors into readable strings
func FormatValidationError(err error) string {
	if ve, ok := err.(validator.ValidationErrors); ok {
		// Just return the first error for simplicity, or loop through all
		for _, e := range ve {
			switch e.Tag() {
			case "required":
				return fmt.Sprintf("Field '%s' is required", e.Field())
			case "email":
				return fmt.Sprintf("Field '%s' must be a valid email", e.Field())
			case "min":
				return fmt.Sprintf("Field '%s' must be at least %s characters long", e.Field(), e.Param())
			case "max":
				return fmt.Sprintf("Field '%s' must not exceed %s characters", e.Field(), e.Param())
			}
		}
	}
	return "Invalid request payload"
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func ComparePassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}