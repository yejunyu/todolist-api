package repository

import (
	"todolist-api/internal/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	// CreateUser  user method
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)

	Create(todo *models.Todo) error
	GetAll(uid uint) ([]models.Todo, error)
	GetById(id uint) (*models.Todo, error)
	Update(id uint) error
	Delete(id uint) error
}
type todoRepository struct {
	db *gorm.DB
}

func (t *todoRepository) CreateUser(user *models.User) error {
	return t.db.Create(user).Error
}
func (t *todoRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := t.db.Where("username=?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (t *todoRepository) Create(todo *models.Todo) error {
	return t.db.Create(todo).Error
}

func (t *todoRepository) GetAll(uid uint) ([]models.Todo, error) {
	var todos []models.Todo
	err := t.db.Where("user_id = ?", uid).Order("created_at desc").Find(&todos).Error
	return todos, err
}

func (t *todoRepository) GetById(id uint) (*models.Todo, error) {
	var todo models.Todo
	err := t.db.First(&todo, id).Error
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t *todoRepository) Update(id uint) error {
	var todo models.Todo

	err := t.db.First(&todo, id).Error
	if err != nil {
		return err
	}
	todo.Status = !todo.Status
	return t.db.Save(todo).Error
}

func (t *todoRepository) Delete(id uint) error {
	return t.db.Delete(&models.Todo{}, id).Error
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db: db}
}
