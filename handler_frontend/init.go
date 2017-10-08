package handler_frontend

import (
	"git.vdo.space/berk/common"
	"gopkg.in/mgo.v2"
)

type FController struct {
	common.Controller
}

func NewFController(MDBSession *mgo.Session, MDB *mgo.Database) *FController {
	mc := &FController{}
	mc.MDB = MDB
	return mc
}
