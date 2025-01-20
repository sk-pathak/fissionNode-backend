package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sk-pathak/fissionNode-backend/internal/app/service"
	"github.com/sk-pathak/fissionNode-backend/internal/db"
)


type NodeHandler struct {
	service *service.NodeService
}

func NewNodeHandler(service *service.NodeService) *NodeHandler {
	return &NodeHandler{service: service}
}

func (h *NodeHandler) RegisterNode(c *gin.Context) {
	var node db.Node
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid input data, " + err.Error()})
		return
	}

	ctx := c.Request.Context()
	if err := h.service.RegisterNode(ctx, &node); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create node, " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, node)
}

func (h *NodeHandler) GetNodeByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	
	ctx := c.Request.Context()

	node, err := h.service.GetNodeByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find node, " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, node)
}

func (h *NodeHandler) GetNodes(c *gin.Context) {
	ctx := c.Request.Context()

	nodes, err := h.service.GetNodes(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch nodes, " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, nodes)
}

func (h *NodeHandler) DeleteNode(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	
	ctx := c.Request.Context()

	err = h.service.DeleteNode(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not find node, " + err.Error()})
		return
	}

    c.JSON(http.StatusOK, gin.H{"message": "Node deleted successfully"})
}

func (h *NodeHandler) UpdateStatus(c *gin.Context) {
    var node db.Node
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid input data, " + err.Error()})
		return
	}
    
    ctx := c.Request.Context()
    if err := h.service.UpdateStatus(ctx, &node); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update node status, " + err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Node status updated successfully"})
}
