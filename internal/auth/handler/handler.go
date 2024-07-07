package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"shop-backend/internal/domain/user"
	"shop-backend/internal/lib/jwt"
)

type AuthHandler struct {
	logger      *slog.Logger
	authService AuthService
}

type AuthService interface {
	Login(ctx context.Context, username, password string) (string, error)
}

func NewHandler(logger *slog.Logger, authService AuthService) *AuthHandler {
	return &AuthHandler{
		logger:      logger,
		authService: authService,
	}
}

func (h *AuthHandler) Register(router *gin.Engine) {
	router.Use()
	router.GET("/login", h.Login)
	router.GET("/signup", h.SignUp)
	router.GET("/test", jwt.VerifyJwt(), h.Test)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var userDTO user.UserDTO

	err := c.BindJSON(&userDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Login(context.Background(), userDTO.Login, userDTO.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var userDTO user.UserDTO

	err := c.BindJSON(&userDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

func (h *AuthHandler) Test(c *gin.Context) {

}
