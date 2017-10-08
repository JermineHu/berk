package handler_backend

import (
	"git.vdo.space/berk/common"
	"gopkg.in/mgo.v2"
)

type BController struct {
	common.Controller
}

func NewBController(MDBSession *mgo.Session, MDB *mgo.Database) *BController {
	mc := &BController{}
	mc.MDB = MDB
	return mc
}
