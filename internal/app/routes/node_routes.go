package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sk-pathak/fissionNode-backend/internal/app/handler"
)

func RegisterNodeRoutes(r *gin.Engine, nodeHandler *handler.NodeHandler) {
	r.POST("/api/nodes", nodeHandler.RegisterNode)
	r.GET("/api/nodes/:id", nodeHandler.GetNodeByID)
	r.GET("/api/nodes", nodeHandler.GetNodes)
	r.DELETE("/api/nodes/:id", nodeHandler.DeleteNode)
	r.PATCH("/api/nodes/status", nodeHandler.UpdateStatus)
}
