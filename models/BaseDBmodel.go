package models

import (
	"gopkg.in/mgo.v2"
	"log"
)

type BaseDBmodel struct {
	session *mgo.Session
	db      *mgo.Database
	c       *mgo.Collection
}

func (this *BaseDBmodel) DBname() string {
	return "contacts"
}

func (this *BaseDBmodel) init() {
	newsession, err := mgo.Dial("localhost:27017")
	if err != nil {
		log.Fatalf("mgo init error :%s", err.Error())
		panic(err)
	}
	this.session = newsession
	this.session.SetMode(mgo.Monotonic, true)
	this.db = this.session.DB(this.DBname())
}
