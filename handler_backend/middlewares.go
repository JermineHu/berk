package handler_backend

import (
	"bufio"
	"bytes"
	. "git.vdo.space/berk/common"
	. "git.vdo.space/berk/consts"
	ml "git.vdo.space/berk/model"
	"git.vdo.space/foy/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"qiniupkg.com/x/errors.v7"
)

type bufferedWriter struct {
	gin.ResponseWriter
	out    *bufio.Writer
	Buffer bytes.Buffer
}

func (g *bufferedWriter) Write(data []byte) (int, error) {
	g.Buffer.Write(data)
	return g.out.Write(data)
}

func GetRequestHeader(header http.Header, key string) string {
	if values, _ := header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}

func AuthorizationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet(ControllerBKey).(*BController)
		token := GetRequestHeader(c.Request.Header, Token)
		if token == "" {
			log.Print("token not found")
			c.Abort()
			panic(model.NewError(AuthErr))
		}
		//uID, err := dm.GetUserIDByTokenWithString(token)
		//if err != nil || len(uID) == 0 {
		//	log.Print("token not found", uID)
		//	c.Abort()
		//	panic(NewError(AuthErr))
		//}
		//u, err := ml.GetUserByID(bson.ObjectIdHex(uID))
		//if err != nil || u == nil {
		//	log.Print("user not found", uID, " ", err)
		//	c.Abort()
		//	panic(NewError(AuthErr))
		//}

		u, err := ml.GetUserByToken(token)
		if err != nil || u == nil {
			c.Abort()
			panic(model.NewError(errors.New("Token is not found ")))
		}

		mc.User = u
		c.Set(ControllerBKey, mc)
	}
}
