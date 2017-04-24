// Copyright 2013 Ardan Studios. All rights reserved.
// Use of baseController source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package baseController implements boilerplate code for all baseControllers.
package controllers

import (
	"github.com/astaxie/beego"
)

//** TYPES

type (
	// BaseController composes all required types and behavior.
	BaseController struct {
		beego.Controller
	}
	
	CommonResp struct {
		Code int `json:"code"`
		Msg string `json:"msg"`
	}
)

func (BaseController *BaseController) retJsonObject(base *BaseController, data interface{})  {
	base.Data["json"] = &data
	base.ServeJSON()
}

func (BaseController *BaseController) retError(base *BaseController, retCode int, retMsg string) {
	beego.Error("error return， code: ", retCode, ", msg: ", retMsg)
	resp := CommonResp{ Code: retCode, Msg: retMsg }
	BaseController.retJsonObject(base, resp)
}
