package handlers

import (
	"net/http"
	"strconv"
	"todolist-api/internal/models"
	"todolist-api/internal/repository"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	repo repository.TodoRepository
}

func NewTodoHandler(t repository.TodoRepository) *TodoHandler {
	return &TodoHandler{repo: t}
}

// CreateTodo godoc
// @Summary      创建新的Todo项目
// @Description  为当前认证用户创建一个新的Todo项目
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        todo           body      CreateTodoInput  true  "Todo信息"
// @Success      201  {object}  models.Todo
// @Failure      400  {object}  map[string]interface{}  "请求参数错误"
// @Failure      401  {object}  map[string]interface{}  "未授权"
// @Failure      500  {object}  map[string]interface{}  "服务器内部错误"
// @Router       /todos [post]
// @Security    BearerAuth
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	var input struct {
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo := models.Todo{Title: input.Title, Status: false, UserId: uid.(uint)}
	if err := h.repo.Create(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

// GetAllTodos godoc
// @Summary      获取用户的所有Todo项目
// @Description  获取当前认证用户的所有Todo项目列表
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Todo
// @Failure      401  {object}  map[string]interface{}  "未授权"
// @Failure      500  {object}  map[string]interface{}  "服务器内部错误"
// @Router       /todos [get]
// @Security    BearerAuth
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	uid, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	var todos, err = h.repo.GetAll(uid.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// GetTodoById godoc
// @Summary      根据ID获取Todo项目
// @Description  根据指定的ID获取特定的Todo项目详情
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id             path      int     true  "Todo ID"
// @Success      200  {object}  models.Todo
// @Failure      400  {object}  map[string]interface{}  "无效的ID格式"
// @Failure      401  {object}  map[string]interface{}  "未授权"
// @Failure      404  {object}  map[string]interface{}  "Todo未找到"
// @Router       /todos/{id} [get]
// @Security    BearerAuth
func (h *TodoHandler) GetTodoById(c *gin.Context) {
	_, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	todo, err := h.repo.GetById(uint(uintId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// UpdateTodo godoc
// @Summary      更新Todo项目状态
// @Description  更新指定ID的Todo项目的完成状态
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id             path      int     true  "Todo ID"
// @Success      200  {object}  models.Todo
// @Failure      400  {object}  map[string]interface{}  "无效的ID格式"
// @Failure      401  {object}  map[string]interface{}  "未授权"
// @Failure      404  {object}  map[string]interface{}  "Todo未找到"
// @Router       /todos/{id} [put]
// @Security    BearerAuth
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	_, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := h.repo.Update(uint(uintId)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	todo, _ := h.repo.GetById(uint(uintId))
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo godoc
// @Summary      删除Todo项目
// @Description  删除指定ID的Todo项目
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id             path      int     true  "Todo ID"
// @Success      204  "删除成功"
// @Failure      400  {object}  map[string]interface{}  "无效的ID格式"
// @Failure      401  {object}  map[string]interface{}  "未授权"
// @Failure      500  {object}  map[string]interface{}  "删除失败"
// @Router       /todos/{id} [delete]
// @Security    BearerAuth
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	_, exists := c.Get("uid")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	id := c.Param("id")
	uintId, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if err := h.repo.Delete(uint(uintId)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete todo"})
		return
	}
	c.Status(http.StatusNoContent)
}

// CreateTodoInput 定义了创建Todo时的输入结构
type CreateTodoInput struct {
	Title string `json:"title" binding:"required" example:"完成项目文档"`
}
