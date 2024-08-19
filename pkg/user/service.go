package user

import (
	"errors"
	"rizkiwhy-todo-app/api/presenter"
	"rizkiwhy-todo-app/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type UserServiceImpl interface {
	RegisterUser(request presenter.RegisterUserRequest) (err error)
	LoginUser(request presenter.LoginUserRequest) (t string, err error)
}

type UserService struct {
	UserRepository *UserRepository
}

func NewService(userRepository *UserRepository) UserServiceImpl {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) RegisterUser(request presenter.RegisterUserRequest) (err error) {
	totalUser := s.UserRepository.CountByFilter(BuildCountFilter(request))

	if totalUser > 0 {
		return fiber.NewError(fiber.StatusConflict, "username or email already taken")
	}

	hashedPassword, err := request.HashPassword()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}

	err = s.UserRepository.Create(BuildCreateRequest(request, hashedPassword))

	return
}

func (s *UserService) LoginUser(request presenter.LoginUserRequest) (t string, err error) {
	user, err := s.UserRepository.GetByEmail(request.Email)
	if err != nil {
		return t, err
	}

	err = user.CompareHashAndPassword(request)
	if err != nil {
		return t, errors.New("invalid credentials")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Name
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	claims["iat"] = time.Now().Unix()

	t, err = token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return t, errors.New("internal server error")
	}

	return
}
