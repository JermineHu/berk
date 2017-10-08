package handler_backend

import (
	. "git.vdo.space/berk/common"
	ml "git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
)

//ArticleCreate 创建Article
func (m *BController) ArticlesCreate(ts []*ml.Article) (*ResultModel, error) {

	rm := ResultModel{}
	for i, _ := range ts {
		ts[i].UserID = m.User.ID
	}
	t, err := ml.NewArticles(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = t
	}

	return &rm, nil
}

//ArticleUpdate 修改Article
func (m *BController) ArticlesUpdate(ts []*ml.Article) (*ResultModel, error) {

	rm := ResultModel{}
	for i, _ := range ts {
		ts[i].UserID = m.User.ID
	}
	ts, err := ml.UpdateArticlesByID(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = ts
	}

	return &rm, nil
}

//ArticleList Article列表
func (m *BController) ArticleList(req *proto.RequestParams) (*ResultModel, error) {

	rm := ResultModel{}

	ts, count, err := ml.FindArticlesByParams(req)

	if err != nil {

		//return nil, NewError(err)

		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		ts = ml.SetTagsForArticle(ts)

		mds, err := ml.ArticlesToDatas(ts)
		if err != nil {
			panic(dm.NewError(err))
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
func (m *BController) ArticleDetailByIDs(ids []*bson.ObjectId) (*ResultModel, error) {

	rm := ResultModel{}

	ts, err := ml.GetArticleByIDs(ids, m.User.ID, ml.ArticleStatusNULL)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		ts = ml.SetTagsForArticle(ts)
		rm.Data = ts
	}

	return &rm, nil
}

//ArticleDetailByIDs Article根据id数组获取Article详情
func (m *BController) ArticleDeleteByIDs(ids []*bson.ObjectId) (*ResultModel, error) {

	rm := ResultModel{}

	d := make(dm.Dungeons)
	d["success"] = ml.DeleteArticlesByIDs(ids, m.User.ID)

	rm.Data = d

	return &rm, nil
}
