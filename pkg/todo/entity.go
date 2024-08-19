package todo

import (
	"rizkiwhy-todo-app/api/presenter"
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Title       string
	Description string
	IsDone      bool
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

type CountFilter struct {
	Email *string
}

func BuildCreateRequest(request presenter.CreateTodoRequest) Todo {
	return Todo{
		ID:          uuid.New(),
		UserID:      request.UserID,
		Title:       request.Title,
		Description: request.Description,
		IsDone:      false,
		CreatedAt:   time.Now(),
	}
}

type TodoFilter struct {
	UserID uuid.UUID
	Name   string
	IsDone []bool
}

func BuildFilter(request presenter.GetTodoFilter) (filter TodoFilter) {
	filter = TodoFilter{UserID: request.UserID, Name: request.Name}

	if request.IsDone != nil {
		filter.IsDone = []bool{*request.IsDone}
	} else {
		filter.IsDone = []bool{false, true}
	}

	return
}

func (t Todo) ToTodoResponse() presenter.TodoResponse {
	return presenter.TodoResponse{
		ID:          t.ID,
		UserID:      t.UserID,
		Title:       t.Title,
		Description: t.Description,
		IsDone:      t.IsDone,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}

}

func (todo *Todo) BuildUpdateRequest(request presenter.UpdateTodoRequest) {
	if request.Title != "" {
		todo.Title = request.Title
	}

	if request.Description != "" {
		todo.Description = request.Description
	}

	if request.IsDone != nil {
		todo.IsDone = *request.IsDone
	}
}

func BuildDeleteRequest(request presenter.DeleteTodoRequest) Todo {
	timeNow := time.Now()

	return Todo{
		ID:        request.ID,
		UserID:    request.UserID,
		DeletedAt: &timeNow,
	}
}

func BuildDetailFilter(request presenter.GetDetailFilter) (todo Todo) {
	todo = Todo{ID: request.ID, UserID: request.UserID}

	return
}
