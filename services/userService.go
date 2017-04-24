package services

import (
	log "github.com/goinggo/tracelog"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"go-shiba/models"
	"go-shiba/middlewares/mongodb"
)

func FindUser(service *Service, userID string) (*models.User, error) {
	log.Startedf(service.UserID, "Find", "userID[%s]", userID)
	
	var queryUser models.User
	callbackFunc := func(collection *mgo.Collection) error {
		queryMap := bson.M{"_id": bson.ObjectIdHex(userID)}
		
		log.Trace(service.UserID, "Find", "MGO : db.user.find(%s).limit(1)", mongodb.ToString(queryMap))
		return collection.Find(queryMap).One(&queryUser)
	}
	
	if err := service.DBAction(service.database, models.UserCollection, callbackFunc); err != nil {
		if err != mgo.ErrNotFound {
			log.CompletedError(err, service.UserID, "Find")
			return nil, err
		}
	}
	
	log.Completedf(service.UserID, "Find", "queryUser%+v", &queryUser)
	return &queryUser, nil
}