package route

import (
	"git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	b "git.vdo.space/berk/handler_backend"
	f "git.vdo.space/berk/handler_frontend"
	//"github.com/brandfolder/gin-gorelic"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

type MeanErrorHandleType int

const (
	MeanErrorHandleTypeHeader MeanErrorHandleType = 1 + iota
	MeanErrorHandelTypeBody
)

type Mean struct {
	Route *gin.Engine
}

func (m Mean) Engine() *gin.Engine {
	m.Route = gin.New()

	// Global middleware
	m.Route.Use(gin.Logger())
	m.Route.Use(common.LoggerWithWriter())
	m.Route.Use(gin.Recovery())

	// new relic monitor
	//gorelic.InitNewrelicAgent("f243fdc54ca4b221bbabef85444e798a6d946335", "Berk", false)
	//m.Route.Use(gorelic.Handler)

	m.Route.Use(InitHandler())
	m.Route.Use(HeaderErrorHandler())
	m.Route.Use(ErrorHandler())

	back := m.Route.Group(BackendGroupRouteModuleName)
	{

		v1 := back.Group(GroupRouteVersion1Key)
		{
			b.LoginHandler(v1)
			if !strings.EqualFold(os.Getenv("IsDebug"), "true") {
				v1.Use(b.AuthorizationHandler())
			}
			b.WrapUserRoutes(v1)
			b.WrapArticleRoutes(v1)
			b.WrapTagRoutes(v1)
			b.WrapFormRoutes(v1)
			b.WrapFeedbackRoutes(v1)
			b.WrapLotteryDrawRoutes(v1)
			b.WrapLotterySettingRoutes(v1)
		}

	}

	front := m.Route.Group(FrontendGroupRouteModuleName)
	{

		v1 := front.Group(GroupRouteVersion1Key)
		{
			f.WrapWXRoutes(v1)
			f.WrapFeedbackRoutes(v1)
			v1.Use(f.DomainHandler())
			f.WrapArticleRoutes(v1)
			f.WrapTagRoutes(v1)
			f.WrapFormRoutes(v1)
			v1.Use(f.AuthorizationHandler())
			f.WrapLotteryDrawRoutes(v1)
			//f.WrapUserRoutes(v1)

		}

	}

	res := m.Route.Group(ResourcesGroupRouteModuleName)
	{

		res.Static("/", os.Getenv(ResourcesPath))

	}

	return m.Route
}
