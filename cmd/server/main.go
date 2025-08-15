package main

import (
	"fmt"
	"log"
	"todolist-api/internal/database"
	"todolist-api/internal/handlers"
	"todolist-api/internal/middleware"
	"todolist-api/internal/repository"
	"todolist-api/internal/routes"
	"todolist-api/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.LoadConfig("configs"); err != nil {
		log.Fatalf("could not load config: %v", err)
	}
	// 需要在初始化路由之前
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("cloud not connect db %v", err)
	}
	// 初始化依赖
	todoRepository := repository.NewTodoRepository(db)
	todoHandler := handlers.NewTodoHandler(todoRepository)
	//r := gin.Default()
	// 注册中间件
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.ErrorHandler())
	// 设置路由
	routes.SetupRoutes(r, todoHandler)

	serverAddr := fmt.Sprintf(":%v", config.Cfg.Server.Port)
	log.Printf("Server is running on port %v", config.Cfg.Server.Port)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
