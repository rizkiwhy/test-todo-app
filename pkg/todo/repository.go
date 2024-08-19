package todo

import (
	"gorm.io/gorm"
)

type TodoRepositoryImpl interface {
	Create(todo Todo) (err error)
	GetByID(filter Todo) (todo *Todo, err error)
	GetByFilter(filter TodoFilter) (todos []Todo, err error)
	Update(todo Todo) (err error)
	Delete(todo Todo) (err error)
}

type TodoRepository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) TodoRepositoryImpl {
	return &TodoRepository{DB: DB}
}

func (r *TodoRepository) Create(todo Todo) (err error) {
	err = r.DB.Create(&todo).Error
	return
}

func (r *TodoRepository) GetByID(filter Todo) (todo *Todo, err error) {
	err = r.DB.Where(&Todo{ID: filter.ID, UserID: filter.UserID}).First(&todo).Error
	return
}

func (r *TodoRepository) GetByFilter(filter TodoFilter) (todos []Todo, err error) {
	filter.Name = "%" + filter.Name + "%"

	err = r.DB.
		Where(r.DB.Where("title LIKE ?", filter.Name).Or("description = ?", filter.Name)).
		Where("is_done IN ?", filter.IsDone).
		Where("user_id = ?", filter.UserID).
		Find(&todos).Error

	return
}

func (r *TodoRepository) Update(todo Todo) (err error) {
	err = r.DB.Save(&todo).Error

	return
}

func (r *TodoRepository) Delete(todo Todo) (err error) {
	err = r.DB.Model(&todo).Updates(Todo{DeletedAt: todo.DeletedAt}).Error

	return
}
