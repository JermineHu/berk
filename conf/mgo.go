// This file "database.go" is created by Lincan Li at 6/22/16.
// Copyright © 2016 - Lincan Li. All rights reserved

package conf

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

var (
	mdb          *mgo.Database
	msession     *mgo.Session
	_MGODatabase string
)

func GetMongo() *mgo.Database {
	if mdb == nil {
		msession, mdb = ConnectMongoDB()
	}

	return mdb
}

func ConnectMongoDB() (*mgo.Session, *mgo.Database) {
	var err error

	s := `mongodb://` + os.Getenv("MONGODB_CONNECTION")

	log.Println("MongoDB Connection String ", s)

	MDBSession, err := mgo.Dial(s)
	if err != nil {
		//panic(err)
		log.Println(err)
	}

	MDB := MDBSession.DB(_MGODatabase)

	return MDBSession, MDB
}

type Auto_Inc_Struct struct {
	Name string `bson:"name"`
	ID   int64  `bson:"id"`
}