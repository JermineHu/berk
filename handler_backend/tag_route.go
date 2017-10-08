package handler_backend

import (
	. "git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	. "git.vdo.space/berk/model"
	"git.vdo.space/foy/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func WrapTagRoutes(g *gin.RouterGroup) {
	TagListHandler(g)
	TagsCreateHandler(g)
	TagsUpdateHandler(g)
	TagDetailByIDsHandler(g)
	TagDeleteByIDsHandler(g)
}

//TagListHandler 根据查询参数查询标签的分页数据
func TagListHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ar := model.RequestQueryParams{}

		err := c.Bind(&ar)
		if err != nil {
			panic(model.NewError(RequestParamsErr))
		}

		result, err := m.TagList(model.RequestParams2ProRequestParams(&ar))
		if err != nil {
			panic(model.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(TagListRoute, hf)
}

//TagsCreateHandler 批量创建标签
func TagsCreateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ts := []*Tag{}

		err := c.Bind(&ts)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.TagsCreate(ts)
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(TagsCreateRoute, hf)
}

//TagsUpdateHandler 标签批量更新
func TagsUpdateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ts := []*Tag{}

		err := c.Bind(&ts)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.TagsUpdate(ts)
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.PUT(TagsUpdateRoute, hf)
}

//TagDetailByIDsHandler 根据标签的id数组获取标签的详情列表信息
func TagDetailByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

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

//TagDeleteByIDsHandler 根据id数组删除对应的数据
func TagDeleteByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.TagDeleteByIDs(ids)
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.DELETE(TagDeleteByIDsRoute, hf)
}
