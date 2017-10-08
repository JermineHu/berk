package handler_frontend

import (
	"bufio"
	"bytes"
	"errors"
	. "git.vdo.space/berk/consts"
	"git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"github.com/gin-gonic/gin"
	"net/http"
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
		mc := c.MustGet(ControllerFKey).(*FController)
		token := GetRequestHeader(c.Request.Header, Token)
		if token == "" {
			c.Abort()
			panic(dm.NewError(errors.New("Token was needed")))
		}

		//uID, err := dm.GetUserIDByTokenWithString(token)
		//if err != nil || len(uID) == 0 {
		//	log.Print("token not found", uID)
		//	c.Abort()
		//	panic(AuthErr)
		//}

		//usr, err := model.GetUserByID(bson.ObjectIdHex(uID))
		//if err != nil || usr == nil {
		//	log.Print("user not found", uID, " ", err)
		//	c.Abort()
		//	panic(AuthErr)
		//}

		usr, err := model.GetUserByToken(token)
		if err != nil || usr == nil {
			c.Abort()
			panic(dm.NewError(errors.New("User not found")))
		}

		mc.User = usr
		c.Set(ControllerFKey, mc)
	}
}

func DomainHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		mc := c.MustGet(ControllerFKey).(*FController)
		name := GetRequestHeader(c.Request.Header, "User-Name")
		if name == "" {
			c.Abort()
			panic(dm.NewError(errors.New("User-Name was needed")))
		}
		usr, err := model.GetUserByName(name)
		if err != nil || usr == nil {
			c.Abort()
			panic(dm.NewError(errors.New("User-Name not found")))
		}

		mc.User = usr
		c.Set(ControllerFKey, mc)
	}
}
