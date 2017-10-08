package handler_backend

import (
	"git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
)

//TagCreate 创建标签
func (m *BController) TagsCreate(ts []*model.Tag) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}
	for i, _ := range ts {
		ts[i].UserID = m.User.ID
	}
	t, err := model.NewTags(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = t
	}

	return &rm, nil
}

//TagUpdate 修改标签
func (m *BController) TagsUpdate(ts []*model.Tag) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}
	for i, _ := range ts {
		ts[i].UserID = m.User.ID
	}
	ts, err := model.UpdateTags(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = ts
	}

	return &rm, nil
}

//TagList 标签列表
func (m *BController) TagList(req *proto.RequestParams) (*dm.ResultModel, error) {

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
func (m *BController) TagDetailByIDs(ids []*bson.ObjectId) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ts, err := model.GetTagByIDs(ids, m.User.ID)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = ts
	}

	return &rm, nil
}

//TagDetailByIDs 标签根据id数组获取标签详情
func (m *BController) TagDeleteByIDs(ids []*bson.ObjectId) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	d := make(dm.Dungeons)
	d["success"] = model.DeleteTagByIDs(ids, m.User.ID)

	rm.Data = d

	return &rm, nil
}
