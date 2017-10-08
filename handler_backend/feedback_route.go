package handler_backend

import (
	. "git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	. "git.vdo.space/berk/model"
	"git.vdo.space/foy/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func WrapFeedbackRoutes(g *gin.RouterGroup) {
	FeedbackListHandler(g)
	FeedbacksCreateHandler(g)
	FeedbacksUpdateHandler(g)
	FeedbackDetailByIDsHandler(g)
	FeedbackDeleteByIDsHandler(g)
}

type FeedbackRequestQueryParams struct {
	model.RequestQueryParams
	Where *Feedback `json:"where"`
}

//FeedbackListHandler 根据查询参数查询的分页数据
func FeedbackListHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ar := FeedbackRequestQueryParams{}

		//b := binding.Default(c.Request.Method, c.)

		err := c.Bind(&ar)
		if err != nil {
			panic(model.NewError(RequestParamsErr))
		}

		rqp := model.RequestQueryParams{}
		if ar.Where != nil {
			rqp.Where = ar.Where
		}
		rqp.EndTime = ar.EndTime
		rqp.IsDesc = ar.IsDesc
		rqp.OffsetHead = ar.OffsetHead
		rqp.OffsetTail = ar.OffsetTail
		rqp.OrderBy = ar.OrderBy
		rqp.StartTime = ar.StartTime
		rqp.UserID = ar.UserID
		rqp.Keywords = ar.Keywords

		result, err := m.FeedbackList(model.RequestParams2ProRequestParams(&rqp))
		if err != nil {
			panic(model.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(FeedbackListRoute, hf)
}

//FeedbacksCreateHandler 批量创建
func FeedbacksCreateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ts := []*Feedback{}

		err := c.Bind(&ts)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.FeedbacksCreate(ts)
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(FeedbacksCreateRoute, hf)
}

//FeedbacksUpdateHandler 批量更新
func FeedbacksUpdateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ts := []*Feedback{}

		err := c.Bind(&ts)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.FeedbacksUpdate(ts)
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.PUT(FeedbacksUpdateRoute, hf)
}

//FeedbackDetailByIDsHandler 根据的id数组获取的详情列表信息
func FeedbackDetailByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.FeedbackDetailByIDs(ids)
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(FeedbackDetailByIDsRoute, hf)
}

//FeedbackDeleteByIDsHandler 根据id数组删除对应的数据
func FeedbackDeleteByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.FeedbackDeleteByIDs(ids)
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.DELETE(FeedbackDeleteByIDsRoute, hf)
}
