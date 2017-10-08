package handler_backend

import (
	"git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
)

//FormCreate 创建标签
func (m *BController) FormsCreate(ts []*model.Form) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}
	for i, _ := range ts {
		ts[i].UserID = m.User.ID
	}
	t, err := model.NewForms(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = t
	}

	return &rm, nil
}

//FormUpdate 修改标签
func (m *BController) FormsUpdate(ts []*model.Form) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}
	for i, _ := range ts {
		ts[i].UserID = m.User.ID
	}
	ts, err := model.UpdateForms(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = ts
	}

	return &rm, nil
}

//FormList 标签列表
func (m *BController) FormList(req *proto.RequestParams) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ts, count, err := model.FindFormsByParams(req, m.User.ID)

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

//FormDetailByIDs 标签根据id数组获取标签详情
func (m *BController) FormDetailByIDs(ids []*bson.ObjectId) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ts, err := model.GetFormByIDs(ids, m.User.ID)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = ts
	}

	return &rm, nil
}

//FormDetailByIDs 标签根据id数组获取标签详情
func (m *BController) FormDeleteByIDs(ids []*bson.ObjectId) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	d := make(dm.Dungeons)
	d["success"] = model.DeleteFormByIDs(ids, m.User.ID)

	rm.Data = d

	return &rm, nil
}
