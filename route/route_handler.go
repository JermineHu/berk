package route

import (
	"git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	b "git.vdo.space/berk/handler_backend"
	f "git.vdo.space/berk/handler_frontend"
	dc "git.vdo.space/berk/model"
	"git.vdo.space/foy/model"
	"git.vdo.space/foy/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

type BodyError struct {
	Success bool                   `json:"success"`
	Error   *model.ErrorMsgDefault `json:"error,omitempty"`
	Data    interface{}            `json:"data,omitempty"`
}

func InitHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("P3P", "CP=CURa ADMa DEVa PSAo PSDo OUR BUS UNI PUR INT DEM STA PRE COM NAV OTC NOI DSP COR")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, No-Cache, X-Requested-With, If-Modified-Since, Pragma, Last-Modified, Cache-Control, Expires, Content-Type, X-E4M-With,Authentication-Token,User-Name, *")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}

		mf := f.NewFController(nil, dc.MG)
		mb := b.NewBController(nil, dc.MG)
		c.Set(ControllerFKey, mf)
		c.Set(ControllerBKey, mb)
		c.Next()
	}
}

func HeaderErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(MeanErrorHandelKey, MeanErrorHandleTypeHeader)
		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				HandleError(c, err)

			} else if err := c.Errors.Last(); err != nil {
				HandleError(c, err)

			}
		}()
		c.Next()
	}
}

func HandleError(c *gin.Context, e interface{}) {
	ht := c.MustGet(MeanErrorHandelKey).(MeanErrorHandleType)

	log.Print(e)
	debug.PrintStack()

	var err *model.ErrorMsgDefault

	switch er := e.(type) {
	case *model.ErrorMsgDefault:
		err = er
	case *gin.Error:
		if er == nil {
			return
		}
		meta := er.Meta.(gin.H)
		err = model.NewResponseError(meta["message"].(string), meta["code"].(int), 500, utils.ErrorTypeInternal)

	default:
		err = common.ServerErr

	}

	if ht == MeanErrorHandleTypeHeader {
		c.JSON(err.HttpCode, gin.H{
			"code":    err.ErrorCode,
			"message": err.Message,
		})

	} else {
		err := BodyError{
			Success: false,
			Error:   err,
		}

		c.JSON(200, err)
	}

	c.Abort()
}

func BodyErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(MeanErrorHandelKey, MeanErrorHandelTypeBody)
		c.Next()
	}
}
