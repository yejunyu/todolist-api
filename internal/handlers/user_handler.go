package handlers

import (
	"errors"
	"net/http"
	"todolist-api/internal/models"
	"todolist-api/internal/repository"
	"todolist-api/internal/services"
	"todolist-api/pkg/ierr"
	"todolist-api/pkg/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	repo        repository.TodoRepository
	authService *services.AuthService
}

// RegisterInput 定义了用户注册时需要绑定的数据
type RegisterInput struct {
	Username string `json:"username" binding:"required,min=4,max=20" example:"johndoe"`
	Password string `json:"password" binding:"required,min=6,max=20" example:"password123"`
}

// LoginInput 定义了用户登录时需要绑定的数据
type LoginInput struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Password string `json:"password" binding:"required" example:"password123"`
}

// RegisterResponse 定义了用户注册成功后的响应结构
type RegisterResponse struct {
	Message string `json:"message" example:"User created successfully"`
}

// LoginResponse 定义了用户登录成功后的响应结构
type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

func NewUserHandler(repo repository.TodoRepository, authService *services.AuthService) *UserHandler {
	return &UserHandler{repo: repo, authService: authService}
}

// Register godoc
// @Summary      用户注册
// @Description  创建新的用户账户
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      RegisterInput  true  "用户注册信息"
// @Success      201   {object}  RegisterResponse
// @Failure      400   {object}  map[string]interface{}  "请求参数错误"
// @Failure      409   {object}  map[string]interface{}  "用户名已存在"
// @Failure      500   {object}  map[string]interface{}  "服务器内部错误"
// @Router       /user/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		// 抛出错误，交由中间件处理
		_ = c.Error(ierr.ErrInvalidInput)
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	if err := h.repo.CreateUser(&user); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			// 抛出特定的业务错误
			_ = c.Error(ierr.ErrUsernameExists)
			return
		}

		// 抛出通用的数据库错误
		_ = c.Error(ierr.ErrSystem)
		return
	}

	// 成功的响应保持不变
	response.Success(c, gin.H{
		"id":         user.ID,
		"username":   user.Username,
		"created_at": user.CreatedAt,
	})
}

// Login godoc
// @Summary      用户登录
// @Description  用户登录并获取JWT令牌
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      LoginInput  true  "用户登录信息"
// @Success      200   {object}  LoginResponse
// @Failure      400   {object}  map[string]interface{}  "请求参数错误"
// @Failure      401   {object}  map[string]interface{}  "无效的凭据"
// @Failure      500   {object}  map[string]interface{}  "服务器内部错误"
// @Router       /user/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		// Gin 的 binding 会自动生成更友好的错误信息
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.repo.GetUserByUsername(input.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response.Success(c, token)
}
