package user

import (
	"errors"

	"github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl interface {
	CountByFilter(filter CountFilter) (result int64)
	Create(user User) (err error)
	GetByEmail(email string) (user *User, err error)
	GetByID(userID uuid.UUID) (user *User, err error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) UserRepositoryImpl {
	return &UserRepository{DB: DB}
}

func (r *UserRepository) CountByFilter(filter CountFilter) (result int64) {
	r.DB.Model(&User{}).Where("email = ?", filter.Email).Count(&result)

	return
}

func (r *UserRepository) Create(user User) (err error) {
	err = r.DB.Create(&user).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		switch mysqlErr.Number {
		case 1062: // MySQL code for duplicate entry
			// Handle duplicate entry
		case 1452: // MySQL code for foreign key violation
		// Add cases for other specific error codes
		case 1364: // MySQL code for required field is missing
		// Add cases for other specific error codes
		default:
			// Handle other errors
			err = errors.New("internal server error")
		}
	}

	return
}

func (r *UserRepository) GetByEmail(email string) (user *User, err error) {
	err = r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("user not found")
		} else {
			err = errors.New("internal server error")
		}
	}

	return
}

func (r *UserRepository) GetByID(userID uuid.UUID) (user *User, err error) {
	err = r.DB.Where("id = ?", userID).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fiber.NewError(fiber.StatusNotFound, "user not found")
		} else {
			err = fiber.NewError(fiber.StatusInternalServerError, ("internal server error"))
		}
	}

	return
}
