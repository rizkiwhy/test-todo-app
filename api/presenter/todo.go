package presenter

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateTodoRequest struct {
	UserID      uuid.UUID `json:"user_id" validate:"required"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
}

type GetDetailFilter struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

type TodoResponse struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

type GetTodoFilter struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	Name   string    `json:"name"`
	IsDone *bool     `json:"is_done"`
}

type UpdateTodoRequest struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	UserID      uuid.UUID `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      *bool     `json:"is_done"`
}

type DeleteTodoRequest struct {
	ID     uuid.UUID `json:"id" validate:"required"`
	UserID uuid.UUID `json:"user_id" validate:"required"`
}

func TodoSuccessResponse(todo TodoResponse) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   todo,
		"error":  nil,
	}
}

func TodosSuccessResponse(todos []TodoResponse) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   todos,
		"error":  nil,
	}
}
