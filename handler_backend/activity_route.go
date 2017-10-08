package handler_backend

import (
	"git.vdo.space/foy/config"
	dm "git.vdo.space/foy/model"
	. "git.vdo.space/foy/utils"
	. "git.vdo.space/berk/consts"
	"git.vdo.space/berk/model"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2/bson"
	"io"
	"os"
	"path"
	"qiniupkg.com/x/errors.v7"
)

func WrapActivityRoutes(g *gin.RouterGroup) {
	ActivityListHandler(g)
	ActivitysCreateHandler(g)
	ActivitysUploadHandler(g)
	ActivitysUpdateHandler(g)
	ActivityDetailByIDsHandler(g)
	ActivityDeleteByIDsHandler(g)
}

type ActivityRequestQueryParams struct {
	dm.RequestQueryParams
	Where ActivityParams `json:"where"`
}

type ActivityParams struct {
	model.Activity
	Status string `json:"status"`
	Type   string `json:"type"`
}

//ActivityListHandler 根据查询参数查询Activity的分页数据
func ActivityListHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ar := ActivityRequestQueryParams{}

		err := c.Bind(&ar)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		at := ar.Where.Activity
		at.Status = model.GetActivityStatusByStr(ar.Where.Status)
		at.Type = model.GetActivityTypeByStr(ar.Where.Status)
		at.Lang = ar.Where.Lang
		at.Region = ar.Where.Region
		at.GroupID = ar.Where.GroupID
		at.ID = ar.Where.ID

		rqp := dm.RequestQueryParams{}
		//rqp.Where = model.Activity2Echo(&at)
		rqp.EndTime = ar.EndTime
		rqp.IsDesc = ar.IsDesc
		rqp.OffsetHead = ar.OffsetHead
		rqp.OffsetTail = ar.OffsetTail
		rqp.OrderBy = ar.OrderBy
		rqp.StartTime = ar.StartTime
		rqp.UserID = ar.UserID
		rqp.Keywords = ar.Keywords

		result, err := m.ActivityList(dm.RequestParams2ProRequestParams(&rqp))
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(ActivityListRoute, hf)
}

//ActivitysCreateHandler 批量创建Activity
func ActivitysCreateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		t := []*ActivityParams{}

		err := c.Bind(&t)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		ts := []*model.Activity{}

		for _, i := range t {

			rqp := i.Activity
			rqp.Type = model.GetActivityTypeByStr(i.Type)
			rqp.Status = model.GetActivityStatusByStr(i.Status)



			ts = append(ts, &rqp)

		}

		result, err := m.ActivitysCreate(ts)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(ActivitysCreateRoute, hf)
}

//ActivitysUpdateHandler Activity批量更新
func ActivitysUpdateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		t := []*ActivityParams{}

		err := c.Bind(&t)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		ts := []*model.Activity{}

		for _, i := range t {

			rqp := i.Activity
			rqp.Type = model.GetActivityTypeByStr(i.Type)
			rqp.Status = model.GetActivityStatusByStr(i.Status)
			ts = append(ts, &rqp)

		}

		result, err := m.ActivitysUpdate(ts)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.PUT(ActivitysUpdateRoute, hf)
}

//ActivityDetailByIDsHandler 根据Activity的id数组获取Activity的详情列表信息
func ActivityDetailByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		result, err := m.ActivityDetailByIDs(ids)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(ActivityDetailByIDsRoute, hf)
}

//ActivityDeleteByIDsHandler 根据id数组删除对应的数据
func ActivityDeleteByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		result, err := m.ActivityDeleteByIDs(ids)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.DELETE(ActivityDeleteByIDsRoute, hf)
}

func ActivitysUpload(c *gin.Context) *dm.ResultModel {

	filename := ""
	ext := ""
	rm := dm.ResultModel{}

	d := make(dm.Dungeons)

	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
		return &rm
	}

	if header != nil {

		filename = uuid.NewV4().String()

		ext = path.Ext(header.Filename)

		if !model.ExtLimit(ext, ImageFileExts) {

			rm.Errors = append(rm.Errors, dm.NewError(errors.New("This file extra must be with .png .jpg .jpeg ! pls try again! ")))
			return &rm

		}

		out, err := os.OpenFile(config.GetResourcePath()+"activities/"+filename+ext, os.O_RDWR|os.O_CREATE, 0644) //打开文件

		if err != nil {
			rm.Errors = append(rm.Errors, dm.NewError(err))
			return &rm
		}

		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			rm.Errors = append(rm.Errors, dm.NewError(err))
			return &rm
		}
	}

	rm.Data = d
	d["file_name"] = filename + ext

	return &rm
}

//ActivitysUploadHandler 批量创建Activity
func ActivitysUploadHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		c.JSON(GenericsSuccessCode, ActivitysUpload(c))

	}
	g.POST(ActivitysUploadRoute, hf)
}
