package model

import (
	"errors"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
)

var MDBVersion *mgo.Collection

var MDB *mgo.Database

func InitMDBVersion(mdb *mgo.Database) {
	MDB = mdb
	MDBVersion = MDB.C("version")
}

type Version struct {
	ID        *bson.ObjectId `bson:"_id" json:"id"`
	Name      string         `bson:"name" json:"name"`
	Version   int64          `bson:"version" json:"version"`
	CreatedAt *time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt *time.Time     `bson:"updated_at" json:"updated_at"`
}

//GetMoodByNum 根据数据的num查询数据
func GetVersionByTableName(name string) (*Version, error) {
	if len(strings.TrimSpace(name)) == 0 {

		return nil, errors.New("The TableName does not a null string,try agin!")
	}
	var ts *Version
	where := bson.M{}
	where["name"] = name
	err := MDBVersion.Find(where).One(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//GetNewVersion 根据表名称增加版本号
func IncreseVersion(name string, mdb *mgo.Database) (*Version, error) {

	now := time.Now()
	update := make(bson.M)
	col := make(bson.M)
	update["$set"] = col
	col[`updated_at`] = &now
	update[`$inc`] = bson.M{"version": 1}

	var v Version
	change := mgo.Change{
		Update:    update,
		ReturnNew: true,
		Upsert:    true,
	}
	_, err := MDBVersion.Find(bson.M{"name": name}).Apply(change, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
