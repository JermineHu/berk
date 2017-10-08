package handler_backend

import (
	. "git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	ml "git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"os"
)

func WrapArticleRoutes(g *gin.RouterGroup) {
	ArticleListHandler(g)
	ArticlesCreateHandler(g)
	ArticlesUploadHandler(g)
	ArticlesUpdateHandler(g)
	ArticleDetailByIDsHandler(g)
	ArticleDeleteByIDsHandler(g)
}

//ArticleListHandler 根据查询参数查询Article的分页数据
func ArticleListHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)
		ar := ml.ArticleRequestQueryParams{}

		err := c.Bind(&ar)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		at := ar.Where.Article
		at.Status = ml.GetArticleStatusByStr(ar.Where.Status)

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

//ArticlesCreateHandler 批量创建Article
func ArticlesCreateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ts := []*ml.ArticleParams{}
		err := c.Bind(&ts)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		al := []*ml.Article{}

		for _, item := range ts {
			x := ml.Article{}
			x.Content = item.Content
			x.Description = item.Description
			x.KeyWords = item.KeyWords
			x.Title = item.Title
			x.CreatedAt=item.CreatedAt
			x.UserID = item.UserID
			x.TagIDs = item.TagIDs
			x.Status = ml.GetArticleStatusByStr(item.Status)
			al = append(al, &x)
		}

		result, err := m.ArticlesCreate(al)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(ArticlesCreateRoute, hf)
}

//ArticlesUpdateHandler Article批量更新
func ArticlesUpdateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ts := []*ml.ArticleParams{}

		err := c.Bind(&ts)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}
		al := []*ml.Article{}

		for _, item := range ts {
			x := ml.Article{}
			x.Content = item.Content
			x.Description = item.Description
			x.KeyWords = item.KeyWords
			x.Title = item.Title
			x.CreatedAt=item.CreatedAt
			x.UserID = item.UserID
			x.TagIDs = item.TagIDs
			x.Status = ml.GetArticleStatusByStr(item.Status)
			x.ID = item.ID
			al = append(al, &x)
		}

		result, err := m.ArticlesUpdate(al)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.PUT(ArticlesUpdateRoute, hf)
}

//ArticleDetailByIDsHandler 根据Article的id数组获取Article的详情列表信息
func ArticleDetailByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

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

//ArticleDeleteByIDsHandler 根据id数组删除对应的数据
func ArticleDeleteByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		result, err := m.ArticleDeleteByIDs(ids)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.DELETE(ArticleDeleteByIDsRoute, hf)
}

//ArticlesUploadHandler 批量创建Article
func ArticlesUploadHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {
		m := c.MustGet(ControllerBKey).(*BController)

		pathstr := "articles/" + m.User.ID.Hex() + "/"

		err := os.MkdirAll(os.Getenv(ResourcesPath)+pathstr, os.ModePerm)

		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, FileUpload(c, pathstr))

	}
	g.POST(ArticlesUploadRoute, hf)
}
