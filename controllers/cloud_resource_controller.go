package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kidus-tiliksew/aqua-crims/application"
	"github.com/kidus-tiliksew/aqua-crims/application/commands"
	"github.com/kidus-tiliksew/aqua-crims/application/queries"
)

type CloudResourceController struct {
	app application.App
}

func NewCloudResourceController(app application.App) *CloudResourceController {
	return &CloudResourceController{app}
}

func (c *CloudResourceController) CloudResourceCreate(ctx *gin.Context) {
	var cloudResource commands.CloudResourceCreate
	if err := ctx.ShouldBind(&cloudResource); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	created, err := c.app.CloudResourceCreate(ctx, cloudResource)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, created)
}

func (c *CloudResourceController) CloudResourceUpdate(ctx *gin.Context) {
	var cloudResource commands.CloudResourceUpdate
	if err := ctx.ShouldBind(&cloudResource); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.app.CloudResourceUpdate(ctx, cloudResource); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, cloudResource)
}

func (c *CloudResourceController) CloudResourceDelete(ctx *gin.Context) {
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

	if err := c.app.CloudResourceDelete(ctx, commands.CloudResourceDelete{ID: int64(id)}); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Cloud Resource deleted successfully"})
}

func (c *CloudResourceController) CloudResourceGet(ctx *gin.Context) {
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

	res, err := c.app.CloudResourceGet(ctx, queries.CloudResourceGet{ID: int64(id)})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, res)
}

func (c *CloudResourceController) CloudResourceFindByCustomer(ctx *gin.Context) {
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

	res, err := c.app.CloudResourceGetByCustomer(ctx, queries.CloudResourceGetByCustomer{CustomerID: int64(id)})
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if res == nil {
		ctx.JSON(404, gin.H{"error": "No resources found for the customer"})
		return
	}

	ctx.JSON(200, res)
}
