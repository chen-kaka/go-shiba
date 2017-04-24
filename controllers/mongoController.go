// Copyright 2013 Ardan Studios. All rights reserved.
// Use of mongoController source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package mongoController implements boilerplate code for all mongoControllers.
package controllers

import (
	"reflect"
	"runtime"

	"fmt"
	"github.com/astaxie/beego/validation"
	"go-shiba/services"
	"go-shiba/middlewares/mongodb"
	log "github.com/goinggo/tracelog"
)

//** TYPES

type (
	// MongoController composes all required types and behavior.
	MongoController struct {
		BaseController
		services.Service
	}
)

//** INTERCEPT FUNCTIONS

// Prepare is called prior to the mongoController method.
func (mongoController *MongoController) Prepare() {
	mongoController.UserID = mongoController.GetString("userID")
	if mongoController.UserID == "" {
		mongoController.UserID = mongoController.GetString(":userID")
	}
	if mongoController.UserID == "" {
		mongoController.UserID = "Unknown"
	}

	if err := mongoController.Service.Prepare(); err != nil {
		log.Errorf(err, mongoController.UserID, "MongoController.Prepare", mongoController.Ctx.Request.URL.Path)
		mongoController.ServeError(err)
		return
	}

	log.Trace(mongoController.UserID, "MongoController.Prepare", "UserID[%s] Path[%s]", mongoController.UserID, mongoController.Ctx.Request.URL.Path)
}

// Finish is called once the mongoController method completes.
func (mongoController *MongoController) Finish() {
	defer func() {
		if mongoController.MongoSession != nil {
			mongodb.CloseSession(mongoController.UserID, mongoController.MongoSession)
			mongoController.MongoSession = nil
		}
	}()

	log.Completedf(mongoController.UserID, "Finish", mongoController.Ctx.Request.URL.Path)
}

//** VALIDATION

// ParseAndValidate will run the params through the validation framework and then
// response with the specified localized or provided message.
func (mongoController *MongoController) ParseAndValidate(params interface{}) bool {
	// This is not working anymore :(
	if err := mongoController.ParseForm(params); err != nil {
		mongoController.ServeError(err)
		return false
	}

	var valid validation.Validation
	ok, err := valid.Valid(params)
	if err != nil {
		mongoController.ServeError(err)
		return false
	}

	if ok == false {
		// Build a map of the Error messages for each field
		messages2 := make(map[string]string)

		val := reflect.ValueOf(params).Elem()
		for i := 0; i < val.NumField(); i++ {
			// Look for an Error tag in the field
			typeField := val.Type().Field(i)
			tag := typeField.Tag
			tagValue := tag.Get("Error")

			// Was there an Error tag
			if tagValue != "" {
				messages2[typeField.Name] = tagValue
			}
		}

		// Build the Error response
		var errors []string
		for _, err := range valid.Errors {
			// Match an Error from the validation framework Errors
			// to a field name we have a mapping for
			message, ok := messages2[err.Field]
			if ok == true {
				// Use a localized message if one exists
				errors = append(errors, message)
				continue
			}

			// No match, so use the message as is
			errors = append(errors, err.Message)
		}

		mongoController.ServeValidationErrors(errors)
		return false
	}

	return true
}

//** EXCEPTIONS

// ServeError prepares and serves an Error exception.
func (mongoController *MongoController) ServeError(err error) {
	mongoController.Data["json"] = struct {
		Error string `json:"Error"`
	}{err.Error()}
	mongoController.Ctx.Output.SetStatus(500)
	mongoController.ServeJSON()
}

// ServeValidationErrors prepares and serves a validation exception.
func (mongoController *MongoController) ServeValidationErrors(Errors []string) {
	mongoController.Data["json"] = struct {
		Errors []string `json:"Errors"`
	}{Errors}
	mongoController.Ctx.Output.SetStatus(409)
	mongoController.ServeJSON()
}

//** CATCHING PANICS

// CatchPanic is used to catch any Panic and log exceptions. Returns a 500 as the response.
func (mongoController *MongoController) CatchPanic(functionName string) {
	if r := recover(); r != nil {
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		log.Warning(mongoController.Service.UserID, functionName, "PANIC Defered [%v] : Stack Trace : %v", r, string(buf))

		mongoController.ServeError(fmt.Errorf("%v", r))
	}
}

//** AJAX SUPPORT

// AjaxResponse returns a standard ajax response.
func (mongoController *MongoController) AjaxResponse(resultCode int, resultString string, data interface{}) {
	response := struct {
		Result       int
		ResultString string
		ResultObject interface{}
	}{
		Result:       resultCode,
		ResultString: resultString,
		ResultObject: data,
	}

	mongoController.Data["json"] = response
	mongoController.ServeJSON()
}
