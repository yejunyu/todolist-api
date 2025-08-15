package repository

import (
	"fmt"
	"log"
	"testing"
	"todolist-api/internal/models"
	"todolist-api/pkg/config"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var repo TodoRepository

func setup() {
	if err := config.LoadConfig("../../configs"); err != nil { // 注意路径
		log.Fatalf("could not load config for test: %v", err)
	}
	cfg := config.Cfg.TestDatabase
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	// 每次测试前都清空并重新迁移表，保证测试环境干净
	db.Migrator().DropTable(&models.Todo{})
	db.AutoMigrate(&models.Todo{})

	repo = NewTodoRepository(db)
}
func TestMain(m *testing.M) {
	setup()
	m.Run()
}
func TestTodoRepository(t *testing.T) {
	t.Run("Create and Get Todo", func(t *testing.T) {
		// 1. Create
		newTodo := &models.Todo{Title: "Test Todo", Status: false}
		err := repo.Create(newTodo)
		assert.NoError(t, err)
		assert.NotZero(t, newTodo.ID)

		// 2. Get By ID
		foundTodo, err := repo.GetById(newTodo.ID)
		assert.NoError(t, err)
		assert.NotNil(t, foundTodo)
		assert.Equal(t, "Test Todo", foundTodo.Title)
		assert.Equal(t, false, foundTodo.Status)

		// 3. Get All
		todos, err := repo.GetAll()
		assert.NoError(t, err)
		assert.Len(t, todos, 1)
		assert.Equal(t, "Test Todo", todos[0].Title)
	})

	t.Run("Update Todo", func(t *testing.T) {
		// 先创建一个
		todo := &models.Todo{Title: "Todo to be updated", Status: false}
		repo.Create(todo)

		// 更新它
		err := repo.Update(todo.ID)
		assert.NoError(t, err)

		// 再次获取并验证
		updatedTodo, err := repo.GetById(todo.ID)
		assert.NoError(t, err)
		//assert.Equal(t, "Updated Title", updatedTodo.Title)
		assert.Equal(t, true, updatedTodo.Status)
	})

	t.Run("Delete Todo", func(t *testing.T) {
		// 先创建一个
		todo := &models.Todo{Title: "Todo to be deleted", Status: false}
		repo.Create(todo)

		// 删除它
		err := repo.Delete(todo.ID)
		assert.NoError(t, err)

		// 尝试获取，应该会失败
		_, err = repo.GetById(todo.ID)
		assert.Error(t, err)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}
