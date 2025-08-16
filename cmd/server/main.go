package main

import (
	"fmt"
	"log"
	"todolist-api/internal/database"
	"todolist-api/internal/handlers"
	"todolist-api/internal/middleware"
	"todolist-api/internal/repository"
	"todolist-api/internal/routes"
	"todolist-api/internal/services"
	"todolist-api/pkg/config"

	_ "todolist-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Todo List API
// @version         1.0
// @description     A simple Todo List API with user authentication and JWT token support
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:4000
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

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
	authService := services.NewAuthService(&config.Cfg.JWT)
	userHandler := handlers.NewUserHandler(todoRepository, authService)
	//r := gin.Default()
	// 注册中间件
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(middleware.ErrorHandler())

	// 设置Swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置路由
	routes.SetupRoutes(r, todoHandler, userHandler, authService)

	serverAddr := fmt.Sprintf(":%v", config.Cfg.Server.Port)
	log.Printf("Server is running on port %v", config.Cfg.Server.Port)
	log.Printf("Swagger documentation available at: http://localhost%v/swagger/index.html", serverAddr)
	if err := r.Run(serverAddr); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
