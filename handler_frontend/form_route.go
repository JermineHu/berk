package handler_frontend

import (
	. "git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	. "git.vdo.space/berk/model"
	"github.com/gin-gonic/gin"
)

func WrapFormRoutes(g *gin.RouterGroup) {
	FormsCreateHandler(g)
}

//FormsCreateHandler 批量创建
func FormsCreateHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		m := c.MustGet(ControllerFKey).(*FController)

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
