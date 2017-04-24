package controllers

import (
	"go-shiba/services"
	"go-shiba/middlewares/exception"
)

type(
	UserController struct {
		MongoController
	}
)

/**
http://localhost:9000/user?id=58fc5435ce01cec8e7f6ec40
 */
func (c *UserController) Get() {
	userId := c.GetString("id")
	if userId == "" {
		c.retError(&c.BaseController, exception.PARAM_MISSING, "id param required.")
		return
	}
	
	userInfo, _ := services.FindUser(&c.Service, userId)
	
	c.retJsonObject(&c.BaseController, userInfo)
}
