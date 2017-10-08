package handler_frontend

import (
	. "git.vdo.space/berk/common"
	ml "git.vdo.space/berk/model"
	"git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"git.vdo.space/foy/utils"
	"github.com/chanxuehong/wechat.v2/open/oauth2"
)

func GetGenderByInt(gender int) ml.Gender {

	switch gender {

	case 2:

		return ml.GENDER_FEMALE

		break

	case 1:

		return ml.GENDER_MALE

		break
	}

	return ml.GENDER_NULL

}

//SetUserInfoByWX 设置微信登陆的用户信息
func (m *FController) SetUserInfoByWX(usr *oauth2.UserInfo) (*ResultModel, error) {

	req := &proto.RequestParams{}

	rm := ResultModel{}

	ts, count, err := ml.FindThirdPartyLoginsByParams(req, ml.LoginType_WECHAT)

	if err != nil {
		return nil, model.NewError(err)
	} else {

		if count == 0 {

			u := ml.User{}

			//avatar:=ml.Avatar{}

			//avatar.

			u.Gender = GetGenderByInt(usr.Sex)
			u.Type = ml.USERTYPE_NORMAL
			u.RealName = usr.Nickname
			u.Nickname = usr.Nickname
			u.Location = usr.Country + "." + usr.Province + "." + usr.City
			//u.Avatars = usr.HeadImageURL
			u.Password = utils.RandomNumberString(6)
			u.Name = usr.UnionId

			usrs, err := ml.NewUsers([]*ml.User{&u})

			if err != nil {
				return nil, model.NewError(err)
			}

			tpl := ml.ThirdPartyLogin{}
			tpl.UserID = usrs[0].ID
			tpl.UnionID = usr.UnionId
			tpl.LoginType = ml.LoginType_WECHAT

			_, err = ml.ThirdPartyLoginsToDatas([]*ml.ThirdPartyLogin{&tpl})

			if err != nil {
				return nil, model.NewError(err)
			}

		}

		pm := PageModelBase{}
		pm.PageData = ts
		pm.Count = count
		rm.Data = pm
	}

	return &rm, nil
}