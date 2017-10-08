package model

import (
	"git.vdo.space/berk/consts"
	"git.vdo.space/foy/model"
	"gopkg.in/mgo.v2"
	"strings"
)

var CUser *mgo.Collection
var CTag *mgo.Collection
var CForm *mgo.Collection
var CArticle *mgo.Collection
var CMenuType *mgo.Collection
var CMenu *mgo.Collection
var CSetting *mgo.Collection
var CAttachment *mgo.Collection
var CAttachmentRelation *mgo.Collection
var CFeedback *mgo.Collection
var CLottery *mgo.Collection
var CLotterySetting *mgo.Collection
var CThirdPartyLogin *mgo.Collection
var CLotteryNum *mgo.Collection
var CActivity *mgo.Collection
var MG *mgo.Database

func InitMDB(mdb *mgo.Database) {
	MG = mdb
	CUser = mdb.C(consts.MDBPerfix + User{}.GetCollectionName())
	CTag = mdb.C(consts.MDBPerfix + Tag{}.GetCollectionName())
	CForm = mdb.C(consts.MDBPerfix + Form{}.GetCollectionName())
	CArticle = mdb.C(consts.MDBPerfix + Article{}.GetCollectionName())
	CMenuType = mdb.C(consts.MDBPerfix + "menu_type")
	CMenu = mdb.C(consts.MDBPerfix + Menu{}.GetCollectionName())
	CSetting = mdb.C(consts.MDBPerfix + Setting{}.GetCollectionName())
	CAttachment = mdb.C(consts.MDBPerfix + Attachment{}.GetCollectionName())
	CAttachmentRelation = mdb.C(consts.MDBPerfix + "attachment_relation")
	CFeedback = mdb.C(consts.MDBPerfix + Feedback{}.GetCollectionName())
	CLottery = mdb.C(consts.MDBPerfix + strings.ToLower(LotteryDraw{}.GetCollectionName()))
	CLotterySetting = mdb.C(consts.MDBPerfix + strings.ToLower(LotterySetting{}.GetCollectionName()))
	CThirdPartyLogin = mdb.C(consts.MDBPerfix + strings.ToLower(ThirdPartyLogin{}.GetCollectionName()))
	CLotteryNum = mdb.C(consts.MDBPerfix + strings.ToLower(LotterySettingNum{}.GetCollectionName()))
	CActivity = mdb.C(consts.MDBPerfix + strings.ToLower(Activity{}.GetCollectionName()))
	model.InitMDBVersion(mdb)
}

//ExtLimit 限制上传的文件类型
func ExtLimit(ext string, exts []string) bool {

	for _, x := range exts {

		if strings.EqualFold(strings.ToLower(ext), x) {

			return true
		}

	}

	return false

}
