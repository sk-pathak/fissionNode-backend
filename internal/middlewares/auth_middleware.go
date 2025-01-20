package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

func AuthMiddleware(store sessions.Store) gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "user-session")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to access session"})
			c.Abort()
			return
		}

		user, ok := session.Values["user"]
		if !ok || user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Set("user", user)

		c.Next()
	}
}
