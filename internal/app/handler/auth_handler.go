package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sk-pathak/fissionNode-backend/internal/app/service"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (h *AuthHandler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the OAuth Demo! Available providers: Google, GitHub, etc.",
	})
}

func (h *AuthHandler) BeginAuth(c *gin.Context) {
	provider := c.Param("provider")
	if provider == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provider not specified"})
		return
	}

	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
	url, err := h.AuthService.BeginAuth(provider, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, url)
}

func (h *AuthHandler) CompleteAuth(c *gin.Context) {
	provider := c.Param("provider")
	if provider == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Provider not specified"})
		return
	}

	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
	user, err := h.AuthService.CompleteAuth(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.AuthService.SaveUserSession(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
	// c.Redirect(http.StatusFound, "/")
}

func (h *AuthHandler) GetUserSession(c *gin.Context) {
	user, ok := h.AuthService.GetUserSession(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No user session found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	err := h.AuthService.Logout(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
