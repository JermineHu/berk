package handler_backend

import (
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"git.vdo.space/berk/model"
	"gopkg.in/mgo.v2/bson"
)

//ActivityCreate 创建Activity
func (m *BController) ActivitysCreate(ts []*model.Activity) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ms, err := model.NewActivitys(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		d, err := model.ActivitysToDatas(ms)
		if err != nil {
			return nil, err
		}
		rm.Data = d
	}

	return &rm, nil
}

//ActivityUpdate 修改Activity
func (m *BController) ActivitysUpdate(imgf []*model.Activity) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ts, err := model.UpdateActivitysByID(imgf)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {

		d, err := model.ActivitysToDatas(ts)
		if err != nil {
			return nil, err
		}
		rm.Data = d
	}

	return &rm, nil
}

//ActivityList Activity列表
func (m *BController) ActivityList(req *proto.RequestParams) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	list, count, err := model.FindActivitysByParams(req)

	if err != nil {

		panic(dm.NewError(err))
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {

		d := []dm.Dungeons{}

		for _, item := range list {

			dg, err := model.ActivityToData(item)

			if err != nil {

				rm.Errors = append(rm.Errors, dm.NewError(err))
			}

			d = append(d, dg)

		}

		pm := dm.PageModelBase{}
		pm.PageData = d
		pm.Count = count
		rm.Data = pm
	}

	return &rm, nil
}

//ActivityDetailByIDs Activity根据id数组获取Activity详情
func (m *BController) ActivityDetailByIDs(ids []*bson.ObjectId) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}
	ts, err := model.GetActivityByIDs(ids)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		d, err := model.ActivitysToDatas(ts)
		if err != nil {
			return nil, err
		}
		rm.Data = d
	}

	return &rm, nil
}

//ActivityDetailByIDs Activity根据id数组获取Activity详情
func (m *BController) ActivityDeleteByIDs(ids []*bson.ObjectId) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	d := make(dm.Dungeons)

	d["success"] = model.DeleteActivitysByIDs(ids)

	rm.Data = d

	return &rm, nil
}
