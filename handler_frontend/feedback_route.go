package handler_frontend

import (
	"errors"
	. "git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	. "git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func WrapFeedbackRoutes(g *gin.RouterGroup) {
	FeedbacksCreateHandler(g)
	FeedbackListHandler(g)
}

type FeedbackCreate struct {
	ID          bson.ObjectId      `bson:"_id" json:"id" `
	VistorID    string             `bson:"vistor_id" json:"vistor_id" binding:"required"`
	VistorName  string             `bson:"vistor_name" json:"vistor_name" binding:"required"`
	Description string             `bson:"description" json:"description" binding:"required"`
	ExtraInfo   interface{}        `bson:"extra_info" json:"extra_info" binding:"required"`
	CreatedAt   *time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt   *time.Time         `bson:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time         `bson:"deleted_at" json:"deleted_at"`
	IsShow      *proto.BooleanType `bson:"is_show" json:"is_show"`
}

//FeedbacksCreateHandler 批量创建
func FeedbacksCreateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		fc := FeedbackCreate{}

		err := c.Bind(&fc)
		if err != nil {
			panic(dm.NewError(err))
		}
		fc.IsShow = &proto.BooleanType{Null: false, Bool: false}

		d, err := GetFeedbackByOpenID(fc.VistorID)

		if err != nil {
			panic(dm.NewError(err))
		}

		if len(d) > 3 {

			panic(dm.NewError(errors.New("One account just submit once!")))
		}

		f := Feedback{}
		f.Description = fc.Description
		f.VistorID = fc.VistorID
		f.VistorName = fc.VistorName
		f.ExtraInfo = fc.ExtraInfo

		result, err := m.FeedbacksCreate([]*Feedback{&f})
		if err != nil {
			panic(err)
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(FeedbacksCreateRoute, hf)
}

//FeedbackListHandler 根据查询参数查询的分页数据
func FeedbackListHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		ar := dm.RequestQueryParams{}

		err := c.Bind(&ar)
		if err != nil {
			panic(RequestParamsErr)
		}

		result, err := m.FeedbackList(dm.RequestParams2ProRequestParams(&ar))
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(FeedbackListRoute, hf)
}
