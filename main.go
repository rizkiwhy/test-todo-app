package main

import (
	"log"
	"rizkiwhy-todo-app/api/handler"
	"rizkiwhy-todo-app/api/router"
	"rizkiwhy-todo-app/database"
	"rizkiwhy-todo-app/pkg/todo"
	"rizkiwhy-todo-app/pkg/user"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	api := app.Group("/api")

	db, err := database.MySQLConnection()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository.(*user.UserRepository))
	userHandler := handler.NewUserHandler(userService.(*user.UserService))
	router.UserRouter(api, userHandler.(*handler.UserHandler))

	todoRepository := todo.NewRepository(db)
	todoService := todo.NewService(todoRepository.(*todo.TodoRepository))
	todoHandler := handler.NewTodoHandler(todoService.(*todo.TodoService))
	router.TodoRouter(api, todoHandler.(*handler.TodoHandler))

	app.Listen(":3000")
}
