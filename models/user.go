package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	UserCollection = "user"
)

type (
	User struct {
		Id 		    bson.ObjectId 	`json:"_id,omitempty" bson:"_id"`
		Username 	string 		    `json:"username,omitempty" bson:"username"`
		CreateAt	time.Time	    `json:"create_at,omitempty" form: "create_at" bson:"create_at"`
	}
)