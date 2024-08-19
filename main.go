package main

import (
	"log"
	"os"
	"rizkiwhy-todo-app/api/handler"
	"rizkiwhy-todo-app/api/router"
	"rizkiwhy-todo-app/database"
	"rizkiwhy-todo-app/pkg/todo"
	"rizkiwhy-todo-app/pkg/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	// Create or open the log file
	logFile, err := os.OpenFile("fiber.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	// Custom logger configuration
	app.Use(logger.New(logger.Config{
		Format:   "[${time}] ${ip} ${method} ${status} ${latency}\n",
		Output:   logFile,
		TimeZone: "UTC",
	}))

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
