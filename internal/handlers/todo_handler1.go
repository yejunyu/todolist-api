package handlers

//
//import (
//	"net/http"
//	"todolist-api/internal/database"
//	"todolist-api/internal/models"
//
//	"github.com/gin-gonic/gin"
//)
//
//func CreateTodo(c *gin.Context) {
//	var input struct {
//		Title string `json:"title" binding:"required"`
//	}
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	todo := models.Todo{Title: input.Title, Status: false}
//	if result := database.DB.Create(&todo); result.Error != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//		return
//	}
//	c.JSON(http.StatusCreated, todo)
//}
//
//func GetAllTodos(c *gin.Context) {
//	var todos []models.Todo
//
//	if result := database.DB.Order("created_at desc").Find(&todos); result.Error != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, todos)
//}
//
//func GetTodoById(c *gin.Context) {
//	var todo models.Todo
//	id := c.Param("id")
//	if result := database.DB.First(&todo, id); result.Error != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, todo)
//}
//
//func UpdateTodo(c *gin.Context) {
//	id := c.Param("id")
//	var todo models.Todo
//
//	if result := database.DB.First(&todo, id); result.Error != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
//		return
//	}
//	todo.Status = !todo.Status
//	if result := database.DB.Model(&todo).Updates(todo); result.Error != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, todo)
//}
//
//func DeleteTodo(c *gin.Context) {
//	var todo models.Todo
//	id := c.Param("id")
//
//	if result := database.DB.First(&todo, id); result.Error != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
//		return
//	}
//	database.DB.Delete(&todo)
//	c.Status(http.StatusNoContent)
//
//}
