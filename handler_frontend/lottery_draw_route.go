package handler_frontend

import (
	dm "git.vdo.space/foy/model"
	. "git.vdo.space/berk/consts"
	"git.vdo.space/berk/model"
	"github.com/gin-gonic/gin"
)

func WrapLotteryDrawRoutes(g *gin.RouterGroup) {
	g.Use(AuthorizationHandler())
	LotteryDrawOutdoorHandler(g)
	//LotteryDrawIndoorHandler(g)
	//LotteryDraw20170610Handler(g)

	//LotteryDrawDetailByIDHamleys20170610Handler(g)
	//LotteryDrawDetailByIDIndoorHandler(g)
	LotteryDrawDetailByIDOutdoorHandler(g)

	//LotteryDrawStatusIndoorHandler(g)
	LotteryDrawStatusOutdoorHandler(g)

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
}

//LotteryDrawOutdoorHandler 室外抽奖
func LotteryDrawOutdoorHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		ld := model.LotteryDraw{}

		ld.LotteryType = model.LotteryTypeVDOOutdoor
		ld.WinningStatus = model.WinningStatusNotWinning
		ld.UserUUID = m.User.ID

		result, err := m.LotteryDrawActivityCreate(ld)

		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.GET(LotteryDrawOutdoorRoute, hf)
}

//LotteryDrawIndoorHandler 室内抽奖
func LotteryDrawIndoorHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)
		ld := model.LotteryDraw{}

		ld.LotteryType = model.LotteryTypeVDOIndoor
		ld.WinningStatus = model.WinningStatusNotWinning
		ld.UserUUID = m.User.ID
		result, err := m.LotteryDrawActivityCreate(ld)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.GET(LotteryDrawIndoorRoute, hf)
}

//LotteryDraw20170610Handler 根据活动类型抽奖
func LotteryDraw20170610Handler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)
		//开始时间戳 1497056400  2017.6.10 9:00
		//结束时间戳 1497175200  2017.6.11 18:00

		//result, err := m.LotteryDrawCreate(model.LotteryTypeHamleys20170610,1497056400,1497175200)
		result, err := m.LotteryDrawCreate(model.LotteryTypeVDO20170710, 1496278800, 1497175200)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.GET(LotteryDraw20170610Route, hf)
}

//LotteryDrawDetailByIDIndoorHandler 根据用户的uuid获取LotteryDraw的indoor详情列表信息
func LotteryDrawDetailByIDIndoorHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		result, err := m.LotteryDrawDetailByUserUUIDAndType(model.LotteryTypeVDOIndoor)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.GET(LotteryDrawDetailByUserUUIDIndoorRoute, hf)
}

//LotteryDrawDetailByIDOutdoorHandler 根据用户的uuid获取LotteryDraw的outdoor详情列表信息
func LotteryDrawDetailByIDOutdoorHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		result, err := m.LotteryDrawDetailByUserUUIDAndType(model.LotteryTypeVDOOutdoor)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.GET(LotteryDrawDetailByUserUUIDOutdoorRoute, hf)
}

//LotteryDrawDetailByIDHamleys20170610Handler 根据用户的uuid获取LotteryDraw的详情列表信息
func LotteryDrawDetailByIDHamleys20170610Handler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		result, err := m.LotteryDrawDetailByUserUUIDAndType(model.LotteryTypeVDO20170710)
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.GET(LotteryDrawDetailByUserUUIDHamleys20170610Route, hf)
}

//LotteryDrawStatusIndoorHandler 获取活动的开始结束状态
func LotteryDrawStatusIndoorHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		result, err := m.GetLotteryDrawStatusIndoor()
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.GET(LotteryDrawIndoorStatusRoute, hf)
}

//LotteryDrawStatusOutdoorHandler 获取活动的开始结束状态
func LotteryDrawStatusOutdoorHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

		result, err := m.GetLotteryDrawStatusOutdoor()
		if err != nil {
			panic(dm.NewError(err))
		}

		c.JSON(GenericsSuccessCode, result)

	}
	g.GET(LotteryDrawOutdoorStatusRoute, hf)
}
