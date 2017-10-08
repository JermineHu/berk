package handler_frontend

import (
	"errors"
	"git.vdo.space/berk/conf"
	. "git.vdo.space/berk/consts"
	"git.vdo.space/foy/model"
	"git.vdo.space/foy/utils"
	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/jssdk"
	o2 "github.com/chanxuehong/wechat.v2/oauth2"
	"github.com/chanxuehong/wechat.v2/open/oauth2"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

func WrapWXRoutes(g *gin.RouterGroup) {
	GetWXConfigHandler(g)
	WXOauth2Handler(g)

}

type WXConfig struct {
	AppId     string `json:"appId"`
	Ticket    string `json:"ticket"`
	Timestamp string `json:"timestamp"`
	NonceStr  string `json:"nonceStr"`
	Signature string `json:"signature"`
}

type ResponesData struct {
	core.Error
	AccessToken string `json:"access_token"`
	OpenID      string `json:"openid"`
}

//GetWXConfigHandler 根据url获取jssdk的配置
func GetWXConfigHandler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {

		url := c.DefaultQuery("url", "")

		if len(strings.TrimSpace(url)) < 10 {

			panic(model.NewError(errors.New("The url is illegal !")))

		}

		wxc := WXConfig{}
		wxc.AppId = conf.GetWXAppID()
		client := core.NewClient(core.NewDefaultAccessTokenServer(wxc.AppId, conf.GetWXAppSecret(), nil), nil)
		ticket, err := jssdk.NewDefaultTicketServer(client).Ticket()

		if err != nil {

			panic(model.NewError(err))
		}

		wxc.Ticket = ticket
		wxc.NonceStr = utils.RandomString(9)
		wxc.Timestamp = strconv.FormatInt(time.Now().Unix(), 10)
		wxc.Signature = jssdk.WXConfigSign(ticket, wxc.NonceStr, wxc.Timestamp, url)

		c.JSON(GenericsSuccessCode, wxc)

	}
	g.GET(GetWXConfigRoute, hf)
}

//WXOauth2Handler The oauth2 of WeChat
func WXOauth2Handler(g *gin.RouterGroup) {
	hf := func(c *gin.Context) {
		m := c.MustGet(ControllerFKey).(*FController)

		code := c.Param("Code")

		if len(strings.TrimSpace(code)) < 10 {

			panic(model.NewError(errors.New("The Code is illegal !")))

		}

		client := o2.Client{}
		client.Endpoint = oauth2.NewEndpoint(conf.GetWXAppID(), conf.GetWXAppSecret())

		token, err := client.ExchangeToken(code)

		if err != nil {
			panic(model.NewError(err))
		}

		usr, err := oauth2.GetUserInfo(token.AccessToken, token.OpenId, "", nil)
		if err != nil {
			panic(model.NewError(err))
		}
		_, err = m.SetUserInfoByWX(usr)

		if err != nil {
			panic(model.NewError(err))
		}

		c.JSON(GenericsSuccessCode, &usr)

	}
	g.GET(WXOauth2Route, hf)
}
