package model

import (
	"go-wss/src/db"
	"gopkg.in/mgo.v2/bson"
)

type Model struct {
	Id bson.ObjectId `json:"_id" bson:"_id"`
}

// collection
var CPlayer = db.DefaultMongoC("player")
