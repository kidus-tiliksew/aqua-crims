package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kidus-tiliksew/aqua-crims/application"
)

type NotificationController struct {
	app application.App
}

func NewNotificationController(app application.App) *NotificationController {
	return &NotificationController{app}
}

func (n *NotificationController) DeleteNotification(ctx *gin.Context) {
	idParam := ctx.Param("id")
	if idParam == "" {
		ctx.JSON(400, gin.H{"error": "id is required"})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "id should be an integer"})
		return
	}

	if err := n.app.NotificationDelete(ctx, int64(id)); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, nil)
}

func (n *NotificationController) DeleteNotificationByUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(400, gin.H{"error": "id is required"})
		return
	}

	if err := n.app.NotificationDeleteByUser(ctx, userID); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(204, nil)
}

func (n *NotificationController) NotificationGetByUser(ctx *gin.Context) {
	userID := ctx.Param("id")
	if userID == "" {
		ctx.JSON(400, gin.H{"error": "id is required"})
		return
	}

	notifications, err := n.app.NotificationGetByUser(ctx, userID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, notifications)
}
