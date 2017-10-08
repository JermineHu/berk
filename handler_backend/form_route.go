package handler_backend

import (
	. "git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	. "git.vdo.space/berk/model"
	"git.vdo.space/foy/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func WrapFormRoutes(g *gin.RouterGroup) {
	FormListHandler(g)
	FormsCreateHandler(g)
	FormsUpdateHandler(g)
	FormDetailByIDsHandler(g)
	FormDeleteByIDsHandler(g)
}

//FormListHandler 根据查询参数查询的分页数据
func FormListHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ar := model.RequestQueryParams{}

		err := c.Bind(&ar)
		if err != nil {
			panic(model.NewError(RequestParamsErr))
		}

		result, err := m.FormList(model.RequestParams2ProRequestParams(&ar))
		if err != nil {
			panic(model.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(FormListRoute, hf)
}

//FormsCreateHandler 批量创建
func FormsCreateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ts := []*Form{}

		err := c.Bind(&ts)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.FormsCreate(ts)
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(FormsCreateRoute, hf)
}

//FormsUpdateHandler 批量更新
func FormsUpdateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ts := []*Form{}

		err := c.Bind(&ts)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.FormsUpdate(ts)
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.PUT(FormsUpdateRoute, hf)
}

//FormDetailByIDsHandler 根据的id数组获取的详情列表信息
func FormDetailByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.FormDetailByIDs(ids)
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(FormDetailByIDsRoute, hf)
}

//FormDeleteByIDsHandler 根据id数组删除对应的数据
func FormDeleteByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.FormDeleteByIDs(ids)
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.DELETE(FormDeleteByIDsRoute, hf)
}
