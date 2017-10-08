package handler_backend

import (
	dm "git.vdo.space/foy/model"
	. "git.vdo.space/foy/utils"
	. "git.vdo.space/berk/consts"
	"git.vdo.space/berk/model"
	"github.com/gin-gonic/gin"
)

func WrapLotteryDrawRoutes(g *gin.RouterGroup) {
	LotteryDrawListHandler(g)
	LotteryDrawsUpdateHandler(g)
	LotteryDrawDetailByIDsHandler(g)
	LotteryDrawsUpdateAwardHandler(g)
	LotteryDrawsCreateHandler(g)
}

type LotteryDrawRequestQueryParams struct {
	dm.RequestQueryParams
	Where LotteryDrawParams `json:"where"`
}

type LotteryDrawParams struct {
	model.LotteryDraw
	WinningStatus string ` json:"winning_status,omitempty"` // 中奖没有
	LotteryStatus string `json:"lottery_status,omitempty"`  //抽过奖没有
	LotteryType   string `json:"lottery_type,omitempty"`    // 抽奖类型
	LotteryLevel  string `json:"lottery_level,omitempty"`   // 抽奖类型
}

//LotteryDrawListHandler 根据查询参数查询LotteryDraw的分页数据
func LotteryDrawListHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ar := LotteryDrawRequestQueryParams{}

		err := c.Bind(&ar)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		at := ar.Where.LotteryDraw
		at.WinningStatus = model.GetWinningStatusByStr(ar.Where.WinningStatus)
		at.LotteryType = model.GetLotteryTypeByStr(ar.Where.LotteryType)
		at.LotteryStatus = model.GetLotteryStatusByStr(ar.Where.LotteryStatus)

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

		result, err := m.LotteryDrawList(dm.RequestParams2ProRequestParamsWithJson(&rqp))
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(LotteryDrawListRoute, hf)
}

//LotteryDrawsCreateHandler 批量创建LotteryDraw
func LotteryDrawsCreateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		t := []*LotteryDrawParams{}

		err := c.Bind(&t)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		ts := []*model.LotteryDraw{}

		for _, i := range t {

			rqp := i.LotteryDraw
			rqp.WinningStatus = model.GetWinningStatusByStr(i.WinningStatus)
			rqp.LotteryType = model.GetLotteryTypeByStr(i.LotteryType)
			rqp.LotteryLevel = model.GetLotteryLevelByStr(i.LotteryLevel)
			ts = append(ts, &rqp)

		}

		result, err := m.LotteryDrawsCreate(ts)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(LotteryDrawsCreateRoute, hf)
}

//LotteryDrawsUpdateHandler LotteryDraw批量更新
func LotteryDrawsUpdateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*BController)

		t := []*LotteryDrawParams{}

		err := c.Bind(&t)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		ts := []*model.LotteryDraw{}

		for _, i := range t {

			rqp := i.LotteryDraw
			rqp.WinningStatus = model.GetWinningStatusByStr(i.WinningStatus)
			rqp.LotteryStatus = model.GetLotteryStatusByStr(i.LotteryStatus)
			rqp.LotteryType = model.GetLotteryTypeByStr(i.LotteryType)
			ts = append(ts, &rqp)

		}

		result, err := m.LotteryDrawsUpdate(ts)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.PUT(LotteryDrawsUpdateRoute, hf)
}

//LotteryDrawsUpdateAwardHandler LotteryDraw批量更新
func LotteryDrawsUpdateAwardHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []int64{}

		err := c.Bind(&ids)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		result, err := m.LotteryDrawsUpdateAward(ids)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.PUT(LotteryDrawsUpdateAwardRoute, hf)
}

//LotteryDrawDetailByIDsHandler 根据LotteryDraw的id数组获取LotteryDraw的详情列表信息
func LotteryDrawDetailByIDsHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerBKey).(*BController)

		ids := []int64{}

		err := c.Bind(&ids)
		if err != nil {
			panic(dm.NewError(RequestParamsErr))
		}

		result, err := m.LotteryDrawDetailByIDs(ids)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.POST(LotteryDrawDetailByIDsRoute, hf)
}
