package handler_backend

import (

	dm "git.vdo.space/foy/model"
	. "git.vdo.space/foy/utils"
	. "git.vdo.space/berk/consts"
	"git.vdo.space/berk/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

func WrapLotterySettingRoutes(g *gin.RouterGroup) {
	LotterySettingListHandler(g)
	LotterySettingsCreateHandler(g)
	LotterySettingsUpdateHandler(g)
	LotterySettingDetailByIDsHandler(g)
	LotterySettingDeleteByIDsHandler(g)
}

type LotterySettingRequestQueryParams struct {
	dm.RequestQueryParams
	Where LotterySettingParams `json:"where"`
}

type LotterySettingParams struct {
	model.LotterySetting
	LotteryType  string `json:"lottery_type"`  // 抽奖类型
	LotteryLevel string `json:"lottery_level"` // 抽奖级别
}

//LotterySettingListHandler 根据查询参数查询LotterySetting的分页数据
func LotterySettingListHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ar := LotterySettingRequestQueryParams{}

		err := c.Bind(&ar)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		at := ar.Where.LotterySetting
		at.LotteryType = model.GetLotteryTypeByStr(ar.Where.LotteryType)
		at.LotteryLevel = model.GetLotteryLevelByStr(ar.Where.LotteryLevel)

		rqp := dm.RequestQueryParams{}
		rqp.Where = at
		rqp.EndTime = ar.EndTime
		rqp.IsDesc = ar.IsDesc
		rqp.OffsetHead = ar.OffsetHead
		rqp.OffsetTail = ar.OffsetTail
		rqp.OrderBy = ar.OrderBy
		rqp.StartTime = ar.StartTime
		rqp.UserID = ar.UserID
		rqp.Keywords = ar.Keywords

		result, err := m.LotterySettingList(dm.RequestParams2ProRequestParamsWithJson(&rqp))
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(LotterySettingListRoute, hf)
}

//LotterySettingsCreateHandler 批量创建LotterySetting
func LotterySettingsCreateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		t := []*LotterySettingParams{}

		err := c.Bind(&t)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		ts := []*model.LotterySetting{}

		for _, i := range t {

			rqp := i.LotterySetting
			rqp.LotteryType = model.GetLotteryTypeByStr(i.LotteryType)
			rqp.LotteryLevel = model.GetLotteryLevelByStr(i.LotteryLevel)

			ts = append(ts, &rqp)

		}

		result, err := m.LotterySettingsCreate(ts)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(LotterySettingsCreateRoute, hf)
}

//LotterySettingsUpdateHandler LotterySetting批量更新
func LotterySettingsUpdateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		t := []*LotterySettingParams{}

		err := c.Bind(&t)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		ts := []*model.LotterySetting{}

		for _, i := range t {

			rqp := i.LotterySetting
			rqp.LotteryType = model.GetLotteryTypeByStr(i.LotteryType)
			rqp.LotteryLevel = model.GetLotteryLevelByStr(i.LotteryLevel)

			ts = append(ts, &rqp)

		}

		result, err := m.LotterySettingsUpdate(ts)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.PUT(LotterySettingsUpdateRoute, hf)
}

//LotterySettingDetailByIDsHandler 根据LotterySetting的id数组获取LotterySetting的详情列表信息
func LotterySettingDetailByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		result, err := m.LotterySettingDetailByIDs(ids)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(LotterySettingDetailByIDsRoute, hf)
}

//LotterySettingDeleteByIDsHandler 根据id数组删除对应的数据
func LotterySettingDeleteByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []*bson.ObjectId{}

		err := c.Bind(&ids)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		result, err := m.LotterySettingDeleteByIDs(ids)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.DELETE(LotterySettingDeleteByIDsRoute, hf)
}
