// Copyright 2013 Ardan Studios. All rights reserved.
// Use of service source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

// Package services implements boilerplate code for all services.
package services

import (
	"go-shiba/middlewares/mongodb"
	"go-shiba/middlewares/exception"
	log "github.com/goinggo/tracelog"
	"gopkg.in/mgo.v2"
	"github.com/astaxie/beego"
)

//** TYPES

type (
	// Service contains common properties for all services.
	Service struct {
		MongoSession *mgo.Session
		UserID       string
		database string
	}
)

func init() {
	// Pull in the configuration.
	//Service.database = beego.AppConfig.String("mongo.database")
	beego.Info("database name is: ", beego.AppConfig.String("mongo.database"))
}


//** PUBLIC FUNCTIONS

// Prepare is called before any controller.
func (service *Service) Prepare() (err error) {
	service.MongoSession, err = mongodb.CopyMonotonicSession(service.UserID)
	if err != nil {
		log.Error(err, service.UserID, "Service.Prepare")
		return err
	}

	return err
}

// Finish is called after the controller.
func (service *Service) Finish() (err error) {
	defer exception.CatchPanic(&err, service.UserID, "Service.Finish")

	if service.MongoSession != nil {
		mongodb.CloseSession(service.UserID, service.MongoSession)
		service.MongoSession = nil
	}

	return err
}

// DBAction executes the MongoDB literal function
func (service *Service) DBAction(databaseName string, collectionName string, dbCall mongodb.DBCall) (err error) {
	return mongodb.Execute(service.UserID, service.MongoSession, databaseName, collectionName, dbCall)
}
