package controllers

import (
	"go-shiba/services"
)

type(
	UserController struct {
		BaseController
	}
)

/**
http://localhost:9000/user?id=58fc5435ce01cec8e7f6ec40
 */
func (c *UserController) Get() {
	userId := c.GetString("id")
	if userId == "" {
		c.retError(&c.BaseController, 1, "param error.")
		return
	}
	
	userInfo, _ := services.FindUser(&c.Service, userId)
	
	c.retJsonObject(&c.BaseController, userInfo)
}
