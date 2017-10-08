package handler_backend

import (
	"git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
)

//FeedbackCreate 创建意见反馈
func (m *BController) FeedbacksCreate(ts []*model.Feedback) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	t, err := model.NewFeedbacks(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = t
	}

	return &rm, nil
}

//FeedbackUpdate 修改意见反馈
func (m *BController) FeedbacksUpdate(ts []*model.Feedback) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ts, err := model.UpdateFeedbacks(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = ts
	}

	return &rm, nil
}

//FeedbackList 意见反馈列表
func (m *BController) FeedbackList(req *proto.RequestParams) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ts, count, err := model.FindFeedbacksByParams(req, false)

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

//FeedbackDetailByIDs 意见反馈根据id数组获取意见反馈详情
func (m *BController) FeedbackDetailByIDs(ids []*bson.ObjectId) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ts, err := model.GetFeedbackByIDs(ids)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = ts
	}

	return &rm, nil
}

//FeedbackDetailByIDs 意见反馈根据id数组获取意见反馈详情
func (m *BController) FeedbackDeleteByIDs(ids []*bson.ObjectId) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	d := make(dm.Dungeons)
	d["success"] = model.DeleteFeedbackByIDs(ids)

	rm.Data = d

	return &rm, nil
}
