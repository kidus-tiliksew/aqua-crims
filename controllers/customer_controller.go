package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kidus-tiliksew/aqua-crims/application"
	"github.com/kidus-tiliksew/aqua-crims/application/commands"
)

type CustomerController struct {
	app application.App
}

func NewCustomerController(app application.App) *CustomerController {
	return &CustomerController{app}
}

func (c *CustomerController) CustomerCreate(ctx *gin.Context) {
	var customer commands.CustomerCreate
	if err := ctx.ShouldBind(&customer); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	created, err := c.app.CustomerCreate(ctx, customer)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, created)
}

func (c *CustomerController) CustomerCreateCloudResources(ctx *gin.Context) {
	var command commands.CustomerCreateCloudResources
	if err := ctx.ShouldBind(&command); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	res, err := c.app.CustomerCreateCloudResources(ctx, command)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201, res)
}
