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

func WrapUserRoutes(g *gin.RouterGroup) {
	UserListHandler(g)
	UsersCreateHandler(g)
	UsersUploadHandler(g)
	UsersUpdateHandler(g)
	UserDetailByIDsHandler(g)
	UserDeleteByIDsHandler(g)
}

type UserRequestQueryParams struct {
	dm.RequestQueryParams
	Where *ml.User `json:"where"`
}

//UserListHandler 根据查询参数查询User的分页数据
func UserListHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

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
		rqp.UserID = ar.UserID
		rqp.Keywords = ar.Keywords

		result, err := m.UserList(dm.RequestParams2ProRequestParams(&rqp))
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(UserListRoute, hf)
}

//UsersCreateHandler 批量创建User
func UsersCreateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ts := []*ml.User{}

		err := c.Bind(&ts)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		result, err := m.UsersCreate(ts)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(UsersCreateRoute, hf)
}

//UsersUpdateHandler User批量更新
func UsersUpdateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ts := []*ml.User{}

		err := c.Bind(&ts)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		result, err := m.UsersUpdate(ts)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.PUT(UsersUpdateRoute, hf)
}

//UserDetailByIDsHandler 根据User的id数组获取User的详情列表信息
func UserDetailByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

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
	g.POST(UserDetailByIDsRoute, hf)
}

//UserDeleteByIDsHandler 根据id数组删除对应的数据
func UserDeleteByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		result, err := m.UserDeleteByIDs(ids)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.DELETE(UserDeleteByIDsRoute, hf)
}



//UsersUploadHandler 批量创建User
func UsersUploadHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		pathstr := "users/" + m.User.ID.Hex() + "/"

		err := os.MkdirAll(os.Getenv(ResourcesPath)+pathstr, os.ModePerm)


		if err != nil {
			panic(err)
		}


		c.JSON(GenericsSuccessCode, FileUpload(c, pathstr))

	}
	g.POST(UsersUploadRoute, hf)
}
