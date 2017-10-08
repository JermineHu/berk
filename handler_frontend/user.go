package handler_frontend

import (
	. "git.vdo.space/berk/common"
	ml "git.vdo.space/berk/model"
	"git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
)

//UserList User列表
func (m *FController) UserList(req *proto.RequestParams) (*ResultModel, error) {

	rm := ResultModel{}

	ts, count, err := ml.FindUsersByParams(req)

	if err != nil {

		//return nil, model.NewError(err)

		rm.Errors = append(rm.Errors, model.NewError(err))
	} else {
		pm := PageModelBase{}
		pm.PageData = ts
		pm.Count = count
		rm.Data = pm
	}

	return &rm, nil
}

//UserDetailByIDs User根据id数组获取User详情
func (m *FController) UserDetailByIDs(ids []*bson.ObjectId) (*ResultModel, error) {

	rm := ResultModel{}

	ts, err := ml.GetUserByIDs(ids)

	if err != nil {
		rm.Errors = append(rm.Errors, model.NewError(err))
	} else {
		rm.Data = ts
	}

	return &rm, nil
}

//UserDetailByName User根据name获取User详情
func (m *FController) UserDetailByName(name string) (*ResultModel, error) {

	rm := ResultModel{}

	ts, err := ml.GetUserByName(name)

	if err != nil {
		rm.Errors = append(rm.Errors, model.NewError(err))
	} else {
		rm.Data = ts
	}

	return &rm, nil
}
