package router

import (
	"rizkiwhy-todo-app/api/handler"
	"rizkiwhy-todo-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TodoRouter(app fiber.Router, h *handler.TodoHandler) {
	todoAPI := app.Group("/todos")
	todoAPI.Post("/", middleware.Protected(), h.Create())
	todoAPI.Get("/", middleware.Protected(), h.GetByFilter())
	todoAPI.Get("/:id", middleware.Protected(), h.GetByID())
	todoAPI.Put("/:id", middleware.Protected(), h.Update())
	todoAPI.Delete("/:id", middleware.Protected(), h.Delete())
}
