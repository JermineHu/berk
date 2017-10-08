// This file "database.go" is created by Lincan Li at 6/22/16.
// Copyright © 2016 - Lincan Li. All rights reserved

package config

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var (
	mdb      *mgo.Database
	msession *mgo.Session
)

var (
	_MGOUserName string
	_MGOPassword string
	_MGOAddress string
	_MGOPort string
	_MGODatabase string
)

func SetMGOUserName(s string) {
	_MGOUserName = s
}

func SetMGOPassword(s string) {
	_MGOPassword = s
}

func SetMGOAddress(s string) {
	_MGOAddress = s
}

func SetMGOPort(s string) {
	_MGOPort = s
}

func SetMGODatabase(s string) {
	_MGODatabase = s
}

func GetMongo() *mgo.Database {
	if mdb == nil {
		msession, mdb = ConnectMongoDB()
	}

	return mdb
}

func GetMongoConnectionString() string {
	if _MGOUserName == `` || _MGOPassword == `` {
		return `mongodb://` + _MGOAddress + `:` + _MGOPort + `/` + _MGODatabase
	}

	return `mongodb://` + _MGOUserName + `:` + _MGOPassword + `@` + _MGOAddress + `:` + _MGOPort + `/` + _MGODatabase
}

func ConnectMongoDB() (*mgo.Session, *mgo.Database) {
	var err error
	s := GetMongoConnectionString()

	log.Println("MongoDB Connection String ", s)

	MDBSession, err := mgo.Dial(s)
	if err != nil {
		panic(err)
	}

	MDB := MDBSession.DB(_MGODatabase)

	return MDBSession, MDB
}

type Auto_Inc_Struct struct {
	Name string `bson:"name"`
	ID   int64  `bson:"id"`
}

const Auto_Inc = "auto_inc"

func GetAutoIncID(name string, mdb *mgo.Database) (*Auto_Inc_Struct, error) {
	var ais Auto_Inc_Struct
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"id": 1}},
		ReturnNew: true,
		Upsert:    true,
	}
	_, err := mdb.C(Auto_Inc).Find(bson.M{"name": name}).Apply(change, &ais)
	if err != nil {
		return nil, err
	}
	return &ais, nil
}

func UpdataAutoIncID(name string, mdb *mgo.Database, id int64) (*Auto_Inc_Struct, error) {
	var ais Auto_Inc_Struct
	change := mgo.Change{
		Update:    bson.M{"id": id},
		ReturnNew: true,
		Upsert:    true,
	}
	_, err := mdb.C(Auto_Inc).Find(bson.M{"name": name}).Apply(change, &ais)
	if err != nil {
		return nil, err
	}
	return &ais, nil
}

// GetAutoID仅仅查询当前自增的数量，如果没有就自动初始化
func GetAutoID(name string, mdb *mgo.Database) (*Auto_Inc_Struct, error) {
	var ais Auto_Inc_Struct
	err := mdb.C(Auto_Inc).Find(bson.M{"name": name}).One(&ais)
	if err != nil {
		if err == mgo.ErrNotFound {
			return GetAutoIncID(name, mdb)
		}
		return nil, err
	}
	return &ais, nil
}

// ## this is a test for submit!