package user

import (
	"rizkiwhy-todo-app/api/presenter"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Name      string
	Password  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type CountFilter struct {
	Email *string
}

func BuildCountFilter(request presenter.RegisterUserRequest) CountFilter {
	return CountFilter{
		Email: &request.Email,
	}
}

func BuildCreateRequest(request presenter.RegisterUserRequest, hashedPassword string) User {
	return User{
		ID:        uuid.New(),
		Email:     request.Email,
		Password:  hashedPassword,
		IsActive:  true,
		CreatedAt: time.Now(),
		Name:      request.Name,
	}
}

func (user User) CompareHashAndPassword(request presenter.LoginUserRequest) (err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return err
	}

	return
}
