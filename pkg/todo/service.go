package todo

import (
	"rizkiwhy-todo-app/api/presenter"

	"github.com/gofiber/fiber/v2"
)

type TodoServiceImpl interface {
	Create(request presenter.CreateTodoRequest) (err error)
	GetByID(filter presenter.GetDetailFilter) (todo *Todo, err error)
	GetByFilter(filter presenter.GetTodoFilter) (todosResponse []presenter.TodoResponse, err error)
	Update(request presenter.UpdateTodoRequest) (err error)
	Delete(request presenter.DeleteTodoRequest) (err error)
}

type TodoService struct {
	TodoRepository *TodoRepository
}

func NewService(todoRepository *TodoRepository) TodoServiceImpl {
	return &TodoService{
		TodoRepository: todoRepository,
	}
}

func (s *TodoService) Create(request presenter.CreateTodoRequest) (err error) {
	return s.TodoRepository.Create(BuildCreateRequest(request))
}

func (s *TodoService) GetByID(filter presenter.GetDetailFilter) (todo *Todo, err error) {
	return s.TodoRepository.GetByID(BuildDetailFilter(filter))
}

func (s *TodoService) GetByFilter(filter presenter.GetTodoFilter) (todosResponse []presenter.TodoResponse, err error) {
	todos, err := s.TodoRepository.GetByFilter(BuildFilter(filter))
	if err != nil {
		return
	}

	for _, todo := range todos {
		todosResponse = append(todosResponse, todo.ToTodoResponse())
	}

	return
}

func (s *TodoService) Update(request presenter.UpdateTodoRequest) (err error) {
	todo, err := s.TodoRepository.GetByID(BuildDetailFilter(presenter.GetDetailFilter{
		ID:     request.ID,
		UserID: request.UserID,
	}))
	if err != nil {
		return
	}
	if todo == nil {
		return fiber.NewError(fiber.StatusNotFound, "todo not found")
	}

	todo.BuildUpdateRequest(request)

	return s.TodoRepository.Update(*todo)
}

func (s *TodoService) Delete(request presenter.DeleteTodoRequest) (err error) {
	return s.TodoRepository.Delete(BuildDeleteRequest(request))
}
