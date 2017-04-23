package controllers

import (
	"go-shiba/services"
)

type UserController struct {
	BaseController
}

func (c *UserController) Get() {
	userInfo, _ := services.FindUser(&c.Service, "58fc5435ce01cec8e7f6ec40")
	
	c.Data["Website"] = userInfo
}
