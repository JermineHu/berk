package handler_frontend

import (
	. "git.vdo.space/berk/common"
	ml "git.vdo.space/berk/model"
	"git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
)

//ArticleList Article列表
func (m *FController) ArticleList(req *proto.RequestParams) (*ResultModel, error) {

	rm := ResultModel{}

	ts, count, err := ml.FindArticlesByParams(req)

	if err != nil {

		//return nil, model.NewError(err)

		rm.Errors = append(rm.Errors, model.NewError(err))
	} else {
		ts = ml.SetTagsForArticle(ts)
		mds, err := ml.ArticlesToDatas(ts)
		if err != nil {
			panic(model.NewError(err))
		}

		for i, item := range mds {
			delete(item, "content")
			mds[i] = item
		}

		pm := PageModelBase{}
		pm.PageData = mds
		pm.Count = count
		rm.Data = pm
	}

	return &rm, nil
}

//ArticleDetailByIDs Article根据id数组获取Article详情
func (m *FController) ArticleDetailByIDs(ids []*bson.ObjectId) (*ResultModel, error) {

	rm := ResultModel{}

	ts, err := ml.GetArticleByIDs(ids, m.User.ID, ml.ArticleStatusPublished)

	if err != nil {
		rm.Errors = append(rm.Errors, model.NewError(err))
	} else {
		ts = ml.SetTagsForArticle(ts)
		rm.Data = ts
	}

	return &rm, nil
}
