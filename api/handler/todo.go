package handler

import (
	"fmt"
	"rizkiwhy-todo-app/api/presenter"
	"rizkiwhy-todo-app/pkg/todo"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TodoHandlerImpl interface {
	Create() fiber.Handler
	GetByID() fiber.Handler
	GetByFilter() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
}

type TodoHandler struct {
	TodoService *todo.TodoService
}

func NewTodoHandler(todoService *todo.TodoService) TodoHandlerImpl {
	return &TodoHandler{
		TodoService: todoService,
	}
}

func (h *TodoHandler) Create() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request presenter.CreateTodoRequest

		if err = c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		request.UserID, err = uuid.Parse(c.Locals("user_id").(string))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		if errs := validator.New().Struct(request); errs != nil {
			for _, vErr := range errs.(validator.ValidationErrors) {
				switch vErr.Tag() {
				case "oneof":
					err = fmt.Errorf("field '%s' must be one of the following values: %s", vErr.Field(), vErr.Param())
				case "required":
					err = fmt.Errorf("field '%s' is required", vErr.Field())
				default:
					err = fmt.Errorf("field '%s' validation failed: %s", vErr.Field(), vErr.Tag())
				}
			}

			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
			}
		}

		if err = h.TodoService.Create(request); err != nil {
			if er, ok := err.(*fiber.Error); ok {
				return c.Status(er.Code).JSON(presenter.GlobalErrorResponse(err))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(presenter.GlobalErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.GlobalSuccessResponse())
	}
}

func (h *TodoHandler) GetByID() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request presenter.GetDetailFilter

		request.ID, err = uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		request.UserID, err = uuid.Parse(c.Locals("user_id").(string))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		if errs := validator.New().Struct(request); errs != nil {
			for _, vErr := range errs.(validator.ValidationErrors) {
				switch vErr.Tag() {
				case "oneof":
					err = fmt.Errorf("field '%s' must be one of the following values: %s", vErr.Field(), vErr.Param())
				case "required":
					err = fmt.Errorf("field '%s' is required", vErr.Field())
				default:
					err = fmt.Errorf("field '%s' validation failed: %s", vErr.Field(), vErr.Tag())
				}
			}

			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
			}
		}

		result, err := h.TodoService.GetByID(request)
		if err != nil {
			if er, ok := err.(*fiber.Error); ok {
				return c.Status(er.Code).JSON(presenter.GlobalErrorResponse(err))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(presenter.GlobalErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.TodoResponse(*result))
	}
}

func (h *TodoHandler) GetByFilter() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request presenter.GetTodoFilter

		request.UserID, err = uuid.Parse(c.Locals("user_id").(string))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		if queryName := c.Query("name"); queryName != "" {
			request.Name = queryName
		}

		if queryIsDone := c.Query("is_done"); queryIsDone != "" {
			parsedIsDone, err := strconv.ParseBool(queryIsDone)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
			}
			request.IsDone = &parsedIsDone
		}

		if errs := validator.New().Struct(request); errs != nil {
			for _, vErr := range errs.(validator.ValidationErrors) {
				switch vErr.Tag() {
				case "oneof":
					err = fmt.Errorf("field '%s' must be one of the following values: %s", vErr.Field(), vErr.Param())
				case "required":
					err = fmt.Errorf("field '%s' is required", vErr.Field())
				default:
					err = fmt.Errorf("field '%s' validation failed: %s", vErr.Field(), vErr.Tag())
				}
			}

			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
			}
		}

		result, err := h.TodoService.GetByFilter(request)
		if err != nil {
			if er, ok := err.(*fiber.Error); ok {
				return c.Status(er.Code).JSON(presenter.GlobalErrorResponse(err))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(presenter.GlobalErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.TodosSuccessResponse(result))
	}
}

func (h *TodoHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request presenter.UpdateTodoRequest

		if err = c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		request.ID, err = uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		request.UserID, err = uuid.Parse(c.Locals("user_id").(string))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		if errs := validator.New().Struct(request); errs != nil {
			for _, vErr := range errs.(validator.ValidationErrors) {
				switch vErr.Tag() {
				case "oneof":
					err = fmt.Errorf("field '%s' must be one of the following values: %s", vErr.Field(), vErr.Param())
				case "required":
					err = fmt.Errorf("field '%s' is required", vErr.Field())
				default:
					err = fmt.Errorf("field '%s' validation failed: %s", vErr.Field(), vErr.Tag())
				}
			}

			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
			}
		}

		err = h.TodoService.Update(request)
		if err != nil {
			if er, ok := err.(*fiber.Error); ok {
				return c.Status(er.Code).JSON(presenter.GlobalErrorResponse(err))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(presenter.GlobalErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.GlobalSuccessResponse())
	}
}

func (h *TodoHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) (err error) {
		var request presenter.DeleteTodoRequest

		request.ID, err = uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		request.UserID, err = uuid.Parse(c.Locals("user_id").(string))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
		}

		if errs := validator.New().Struct(request); errs != nil {
			for _, vErr := range errs.(validator.ValidationErrors) {
				switch vErr.Tag() {
				case "oneof":
					err = fmt.Errorf("field '%s' must be one of the following values: %s", vErr.Field(), vErr.Param())
				case "required":
					err = fmt.Errorf("field '%s' is required", vErr.Field())
				default:
					err = fmt.Errorf("field '%s' validation failed: %s", vErr.Field(), vErr.Tag())
				}
			}

			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(presenter.GlobalErrorResponse(err))
			}
		}

		err = h.TodoService.Delete(request)
		if err != nil {
			if er, ok := err.(*fiber.Error); ok {
				return c.Status(er.Code).JSON(presenter.GlobalErrorResponse(err))
			}

			return c.Status(fiber.StatusInternalServerError).JSON(presenter.GlobalErrorResponse(err))
		}

		return c.Status(fiber.StatusOK).JSON(presenter.GlobalSuccessResponse())
	}
}
