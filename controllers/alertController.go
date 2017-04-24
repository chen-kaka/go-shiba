package controllers

import (
	"go-shiba/middlewares/wechatalert"
	"go-shiba/middlewares/exception"
)

type(
	AlertController struct {
		BaseController
	}
)

/**
http://localhost:9000/alert?id=10007
 */
func (c *AlertController) Get() {
	wechatId := c.GetString("id")
	if wechatId == "" {
		c.retError(&c.BaseController, exception.PARAM_MISSING, "id param required.")
		return
	}
	
	wechatalert.Alert(wechatId, "hello kaka", "", "", "")
	
	c.retError(&c.BaseController, exception.SUCCESS, "succ")
}
