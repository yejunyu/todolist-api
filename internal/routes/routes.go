// 文件路径: internal/routes/routes.go
package routes

import (
	"todolist-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置所有应用的路由
func SetupRoutes(router *gin.Engine, handlers *handlers.TodoHandler) {
	// 创建一个路由组 /api/v1
	api := router.Group("/api/v1")

	// 为 todos 创建一个子路由组
	todoRoutes := api.Group("/todos")

	todoRoutes.POST("", handlers.CreateTodo)
	todoRoutes.GET("", handlers.GetAllTodos)
	todoRoutes.GET("/:id", handlers.GetTodoById)
	todoRoutes.PUT("/:id", handlers.UpdateTodo)
	todoRoutes.DELETE("/:id", handlers.DeleteTodo)

}
