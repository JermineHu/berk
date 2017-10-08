package handler_frontend

import (
	. "git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	ml "git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func WrapArticleRoutes(g *gin.RouterGroup) {
	ArticleListHandler(g)
	ArticleDetailByIDsHandler(g)
}

//ArticleListHandler 根据查询参数查询Article的分页数据
func ArticleListHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		ar := ml.ArticleRequestQueryParams{}

		err := c.Bind(&ar)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		at := ar.Where.Article
		at.Status = ml.ArticleStatusPublished

		rqp := dm.RequestQueryParams{}
		rqp.Where = at
		rqp.EndTime = ar.EndTime
		rqp.IsDesc = ar.IsDesc
		rqp.OffsetHead = ar.OffsetHead
		rqp.OffsetTail = ar.OffsetTail
		rqp.OrderBy = ar.OrderBy
		rqp.StartTime = ar.StartTime
		rqp.UserUUID = m.User.ID.Hex()
		rqp.Keywords = ar.Keywords

		result, err := m.ArticleList(dm.RequestParams2ProRequestParams(&rqp))
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(ArticleListRoute, hf)
}

//ArticleDetailByIDsHandler 根据Article的id数组获取Article的详情列表信息
func ArticleDetailByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		result, err := m.ArticleDetailByIDs(ids)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(ArticleDetailByIDsRoute, hf)
}
