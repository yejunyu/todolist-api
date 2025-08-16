// 文件路径: internal/routes/routes.go
package routes

import (
	"todolist-api/internal/handlers"
	"todolist-api/internal/middleware"
	"todolist-api/internal/services"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置所有应用的路由
func SetupRoutes(router *gin.Engine, todoHandler *handlers.TodoHandler,
	userHandler *handlers.UserHandler, service *services.AuthService) {
	// 创建一个路由组 /api/v1
	api := router.Group("/api/v1")
	public := api.Group("/user")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", userHandler.Login)
	}

	// 受保护的路由
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware(service))
	{
		// 为 todos 创建一个子路由组
		todoRoutes := protected.Group("/todos")
		{
			todoRoutes.POST("", todoHandler.CreateTodo)
			todoRoutes.GET("", todoHandler.GetAllTodos)
			todoRoutes.GET("/:id", todoHandler.GetTodoById)
			todoRoutes.PUT("/:id", todoHandler.UpdateTodo)
			todoRoutes.DELETE("/:id", todoHandler.DeleteTodo)
		}
	}
}
