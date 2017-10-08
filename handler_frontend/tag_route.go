package handler_frontend

import (
	. "git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	"git.vdo.space/foy/model"
	ml "git.vdo.space/berk/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func WrapTagRoutes(g *gin.RouterGroup) {
	TagListHandler(g)
	TagDetailByIDsHandler(g)

}

//TagListHandler 根据查询参数查询标签的分页数据
func TagListHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)



		ar := ml.TagRequestQueryParams{}

		err := c.Bind(&ar)
		if err != nil {
			panic(model.NewError(RequestParamsErr))
		}

		rqp := model.RequestQueryParams{}
		rqp.Where = ar.Where
		rqp.EndTime = ar.EndTime
		rqp.IsDesc = ar.IsDesc
		rqp.OffsetHead = ar.OffsetHead
		rqp.OffsetTail = ar.OffsetTail
		rqp.OrderBy = ar.OrderBy
		rqp.StartTime = ar.StartTime
		rqp.UserUUID = m.User.ID.Hex()
		rqp.Keywords = ar.Keywords

		result, err := m.TagList(model.RequestParams2ProRequestParams(&rqp))
		if err != nil {
			panic(model.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(TagListRoute, hf)
}

//TagDetailByIDsHandler 根据标签的id数组获取标签的详情列表信息
func TagDetailByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.TagDetailByIDs(ids)
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(TagDetailByIDsRoute, hf)
}
