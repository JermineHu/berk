package handler_backend

import (
	. "git.vdo.space/berk/common"
	ml "git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/utils"
)

// Login 方法: 登陆方法, 用户通过传入的 用户名(可为邮箱和手机号) 和 密码进行登陆
func (m *BController) Login(username, password string) (dm.Dungeons, error) {
	if utils.IsEmail(username) {
		return m.LoginByEmail(username, password)
	}

	if utils.IsMobile(username) {
		return m.LoginByPhoneNumber(username, password)
	}

	return m.LoginByName(username, password)

}

// LoginByEmail 方法: 通过 邮箱 和 密码 进行用户登录, 同时, 登陆成功会保存用户的DeviceToken.
func (m *BController) LoginByEmail(email, password string) (dm.Dungeons, error) {

	u, err := ml.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, InvalidEmailAdressErr
	}

	if u.Type != ml.USERTYPE_ADMINISTRATOR {
		return nil, AuthenticationErr
	}

	if !utils.IsPassword(password) {
		return nil, InvalidPasswordErr
	}

	fpwd := utils.GetFinalPwdBySalt(password, u.Salt)

	if string(fpwd) != u.Password {
		return nil, InvalidPasswordErr
	}
	//_, err = dm.InsertUserTokenWithIDString(u.ID.Hex(), u.Token)
	//if err != nil {
	//	return nil, err
	//}

	uData, err := ml.UserToData(u)
	if err != nil {
		return nil, err
	}

	return uData, nil
}

// LoginByEmail 方法: 通过 手机号 和 密码 进行用户登录, 同时, 登陆成功会保存用户的DeviceToken.
func (m *BController) LoginByPhoneNumber(phoneNumber, password string) (dm.Dungeons, error) {

	u, err := ml.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, InvalidPhoneNumberErr
	}

	if u.Type != ml.USERTYPE_ADMINISTRATOR {
		return nil, AuthenticationErr
	}
	if !utils.IsPassword(password) {
		return nil, InvalidPasswordErr
	}

	fpwd := utils.GetFinalPwdBySalt(password, u.Salt)

	if string(fpwd) != u.Password {
		return nil, InvalidPasswordErr
	}
	//
	//_, err = dm.InsertDeviceToken(u.ID, dToken)
	//if err != nil {
	//	return nil, err
	//}

	//_, err = dm.InsertUserTokenWithIDString(u.ID.Hex(), u.Token)
	//if err != nil {
	//	return nil, err
	//}

	uData, err := ml.UserToData(u)
	if err != nil {
		return nil, err
	}

	return uData, nil
}

// LoginByName 方法: 通过 手机号 和 密码 进行用户登录, 同时, 登陆成功会保存用户的DeviceToken.
func (m *BController) LoginByName(name, password string) (dm.Dungeons, error) {

	u, err := ml.GetUserByName(name)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, UserNotFoundErr
	}

	if u.Type != ml.USERTYPE_ADMINISTRATOR {
		return nil, AuthenticationErr
	}
	if !utils.IsPassword(password) {
		return nil, InvalidPasswordErr
	}

	fpwd := utils.GetFinalPwdBySalt(password, u.Salt)

	if string(fpwd) != u.Password {
		return nil, InvalidPasswordErr
	}
	//
	//_, err = dm.InsertDeviceToken(u.ID, dToken)
	//if err != nil {
	//	return nil, err
	//}

	//_, err = dm.InsertUserTokenWithIDString(u.ID.Hex(), u.Token)
	//if err != nil {
	//	return nil, err
	//}

	uData, err := ml.UserToData(u)
	if err != nil {
		return nil, err
	}

	return uData, nil
}
