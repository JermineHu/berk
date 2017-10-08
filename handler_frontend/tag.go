package handler_frontend

import (
	"git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
)

//TagList 标签列表
func (m *FController) TagList(req *proto.RequestParams) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ts, count, err := model.FindTagsByParams(req, m.User.ID)

	if err != nil {

		//return nil, dm.NewError(err)

		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		pm := dm.PageModelBase{}
		pm.PageData = ts
		pm.Count = count
		rm.Data = pm
	}

	return &rm, nil
}

//TagDetailByIDs 标签根据id数组获取标签详情
func (m *FController) TagDetailByIDs(ids []*bson.ObjectId) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ts, err := model.GetTagByIDs(ids, m.User.ID)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = ts
	}

	return &rm, nil
}
