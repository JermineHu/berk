package handler_backend

import (
	. "git.vdo.space/berk/common"
	ml "git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
)

//UserCreate 创建User
func (m *BController) UsersCreate(ts []*ml.User) (*ResultModel, error) {

	rm := ResultModel{}

	t, err := ml.NewUsers(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = t
	}

	return &rm, nil
}

//UserUpdate 修改User
func (m *BController) UsersUpdate(ts []*ml.User) (*ResultModel, error) {

	rm := ResultModel{}

	ts, err := ml.UpdateUsersByID(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = ts
	}

	return &rm, nil
}

//UserList User列表
func (m *BController) UserList(req *proto.RequestParams) (*ResultModel, error) {

	rm := ResultModel{}

	ts, count, err := ml.FindUsersByParams(req)

	if err != nil {

		//return nil, dm.NewError(err)

		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		pm := PageModelBase{}
		pm.PageData = ts
		pm.Count = count
		rm.Data = pm
	}

	return &rm, nil
}

//UserDetailByIDs User根据id数组获取User详情
func (m *BController) UserDetailByIDs(ids []*bson.ObjectId) (*ResultModel, error) {

	rm := ResultModel{}

	ts, err := ml.GetUserByIDs(ids)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = ts
	}

	return &rm, nil
}

//UserDetailByIDs User根据id数组获取User详情
func (m *BController) UserDeleteByIDs(ids []*bson.ObjectId) (*ResultModel, error) {

	rm := ResultModel{}

	d := make(dm.Dungeons)
	d["success"] = ml.DeleteUsersByIDs(ids)

	rm.Data = d

	return &rm, nil
}
