package handler_frontend

import (
	"git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
)

//FeedbackCreate 创建标签
func (m *FController) FeedbacksCreate(ts []*model.Feedback) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	t, err := model.NewFeedbacks(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = t
	}

	return &rm, nil
}

//FeedbackList 意见反馈列表
func (m *FController) FeedbackList(req *proto.RequestParams) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ts, count, err := model.FindFeedbacksByParams(req, true)

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
