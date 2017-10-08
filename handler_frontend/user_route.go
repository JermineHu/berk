package handler_frontend

import (
	. "git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	ml "git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func WrapUserRoutes(g *gin.RouterGroup) {
	UserListHandler(g)
	UserDetailByIDsHandler(g)
}

type UserRequestQueryParams struct {
	dm.RequestQueryParams
	Where *ml.User `json:"where"`
}

//UserListHandler 根据查询参数查询User的分页数据
func UserListHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		ar := UserRequestQueryParams{}

		err := c.Bind(&ar)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		rqp := dm.RequestQueryParams{}
		if ar.Where != nil {
			rqp.Where = ar.Where
		}
		rqp.EndTime = ar.EndTime
		rqp.IsDesc = ar.IsDesc
		rqp.OffsetHead = ar.OffsetHead
		rqp.OffsetTail = ar.OffsetTail
		rqp.OrderBy = ar.OrderBy
		rqp.StartTime = ar.StartTime
		rqp.UserUUID = m.User.ID.Hex()
		rqp.Keywords = ar.Keywords

		result, err := m.UserList(dm.RequestParams2ProRequestParams(&rqp))
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(UserListRoute, hf)
}

//UserDetailByIDsHandler 根据User的id数组获取User的详情列表信息
func UserDetailByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		result, err := m.UserDetailByIDs(ids)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.GET(UserDetailByIDsRoute, hf)
}
