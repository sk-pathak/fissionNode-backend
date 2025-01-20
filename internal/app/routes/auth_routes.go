package routes

import (
	"github.com/gin-gonic/gin"
	handler "github.com/sk-pathak/fissionNode-backend/internal/app/handler"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler *handler.AuthHandler) {
	r.GET("/", authHandler.Home)
	r.GET("/auth/:provider", authHandler.BeginAuth)
	r.GET("/auth/:provider/complete", authHandler.CompleteAuth)
	r.GET("/logout", authHandler.Logout)
	r.GET("/auth/session", authHandler.GetUserSession)
}
