package handler_backend

import (
	. "git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	"git.vdo.space/foy/model"
	"github.com/gin-gonic/gin"
)

type Login struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func WrapAccountRoutes(g *gin.RouterGroup) {
	//CreateMerchantAccountHandler(g)
	//BatchSetMerchantAccountModeHandler(g)

}

//LoginHandler 后台登录
func LoginHandler(g *gin.RouterGroup) {
	LoginEndPoint := func(c *gin.Context) {
		m := c.MustGet(ControllerBKey).(*BController)

		var login Login
		if err := c.BindJSON(&login); err != nil {
			panic(RequestParamsErr)
		}

		user, err := m.Login(login.UserName, login.Password)
		if err != nil {
			panic(model.NewError(err))
		}

		c.JSON(GenericsSuccessCode, user)
	}
	g.POST(LoginRoute, LoginEndPoint)
}

//
////CreateMerchantAccountHandler 批量创建商户账号
//func CreateMerchantAccountHandler(g *gin.RouterGroup) {
//	hf := func(c *gin.Context) {
//
//		m := c.MustGet(ControllerKey).(*Controller)
//
//		abs := []*model.AccountBaseInfo{}
//
//		err := c.Bind(&abs)
//		if err != nil {
//			panic(RequestParamsErr)
//		}
//
//		result, err := m.CreateMerchantAccount(abs)
//		if err != nil {
//			panic(err)
//		}
//
//		c.JSON(GenericsSuccessCode, result)
//
//	}
//	g.POST(BatchAddMerchantsRoute, hf)
//}
//
////BatchSetMerchantAccountModeHandler 批量设置为商户账号
//func BatchSetMerchantAccountModeHandler(g *gin.RouterGroup) {
//	hf := func(c *gin.Context) {
//
//		m := c.MustGet(ControllerKey).(*Controller)
//
//		ids := []int64{}
//
//		err := c.Bind(&ids)
//		if err != nil {
//			panic(RequestParamsErr)
//		}
//
//		result, err := m.PatchUserTypeWithMerchantByIDs(ids)
//		if err != nil {
//			panic(err)
//		}
//
//		c.JSON(GenericsSuccessCode, result)
//
//	}
//	g.PUT(AccountBatchSetMerchantModeRoute, hf)
//}
