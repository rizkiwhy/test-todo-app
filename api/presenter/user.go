package presenter

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

func (request RegisterUserRequest) HashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	return string(bytes), err
}

func UserLoginSuccessResponse(t string) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   t,
		"error":  nil,
	}
}
