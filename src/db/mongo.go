package db

import (
	"fmt"
	"go-wss/src/config"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

type Collection struct {
	mgo.Collection
}

func MongoSession() *mgo.Session {
	// mongo回话
	mongoAddr := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/?authSource=admin",
		config.System.Mongo.USERNAME, config.System.Mongo.PASSWORD, config.System.Mongo.HOST, config.System.Mongo.PORT,
	)
	if session == nil {
		session, _ = mgo.Dial(mongoAddr)
	}
	return session
}

func MongoDB(name string) *mgo.Database {
	// mongo回话
	db := MongoSession().DB(name)
	return db
}

func DefaultMongoDB() *mgo.Database {
	// 默认的mongo数据库
	return MongoDB(config.System.Mongo.DATABASE)
}

func MongoC(dbName string, cName string) *mgo.Collection {
	// 选择mongo集合
	c := MongoDB(dbName).C(cName)
	return c
}

func DefaultMongoC(cName string) *mgo.Collection {
	// 默认的mongo集合(配置表里的那个库)
	c := DefaultMongoDB().C(cName)
	return c
}
