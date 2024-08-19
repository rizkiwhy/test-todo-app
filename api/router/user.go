package router

import (
	"rizkiwhy-todo-app/api/handler"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, h *handler.UserHandler) {
	userAPI := app.Group("/users")
	userAPI.Post("/register", h.RegisterUser())
	userAPI.Post("/login", h.Login())
}
