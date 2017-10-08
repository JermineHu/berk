package model

import (
	"encoding/json"
	"errors"
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"git.vdo.space/foy/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
	"time"
)

type Gender int

const (
	GENDER_NULL   Gender = 0
	GENDER_FEMALE Gender = 1
	GENDER_MALE   Gender = 2
)

type UserType int

const (
	USERTYPE_NULL          UserType = 0
	USERTYPE_NORMAL        UserType = 1
	USERTYPE_PUP           UserType = 2
	USERTYPE_OPERATIONS    UserType = 3
	USERTYPE_ADMINISTRATOR UserType = 4
	USERTYPE_MERCHANT      UserType = 5
)

type UserStatus int

const (
	USERSTATUS_NULL        UserStatus = 0
	USERSTATUS_ACTIVATED   UserStatus = 1
	USERSTATUS_DEACTIVATED UserStatus = 2
	USERSTATUS_CLOSED      UserStatus = 3
)

type User struct {
	ID           *bson.ObjectId `bson:"_id" json:"id"`
	Name         string         `bson:"name" json:"name"`
	Gender       Gender         `bson:"gender" json:"gender"`
	Nickname     string         `bson:"nickname" json:"nickname"`
	RealName     string         `bson:"real_name" json:"real_name"`
	Birthday     string         `bson:"birthday" json:"birthday"`
	Location     string         `bson:"location" json:"location"`
	PhoneNumber  string         `bson:"phone_number" json:"phone_number"`
	Password     string         `bson:"password" json:"password"`
	Email        string         `bson:"email" json:"email"`
	Salt         string         `bson:"salt" json:"salt"`
	Token        string         `bson:"token" json:"token"`
	Type         UserType       `bson:"type" json:"type"`
	Status       UserStatus     `bson:"status" json:"status"`
	Avatars      []*Avatar      `bson:"avatars" json:"avatars"`
	ContactInfos []*ContactInfo `bson:"contact_infos" json:"contact_infos"`
	CreatedAt    *time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt    *time.Time     `bson:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time     `bson:"deleted_at" json:"deleted_at"`
}

func (v User) GetCollectionName() string {
	return strings.ToLower(reflect.TypeOf(v).Name())
}

func UserToData(md *User) (dm.Dungeons, error) {
	d := make(dm.Dungeons)
	d[`id`] = md.ID
	d[`name`] = md.Name
	d[`gender`] = md.Gender
	d[`nickname`] = md.Nickname
	d[`real_name`] = md.RealName
	d[`birthday`] = md.Birthday
	d[`location`] = md.Location
	d[`phone_number`] = md.PhoneNumber
	d[`email`] = md.Email
	d[`token`] = md.Token
	d[`type`] = md.Type
	d[`status`] = md.Status
	d[`created_at`] = md.CreatedAt
	d[`updated_at`] = md.UpdatedAt
	return d, nil
}

func UsersToDatas(mds []*User) ([]dm.Dungeons, error) {
	dl := []dm.Dungeons{}
	for _, x := range mds {
		d, err := UserToData(x)
		if err != nil {
			return nil, err
		}
		dl = append(dl, d)
	}

	return dl, nil
}

func toCheckThisUser(t *User) error {

	t.Name = strings.TrimSpace(t.Name)
	t.Email = strings.TrimSpace(t.Email)
	t.PhoneNumber = strings.TrimSpace(t.PhoneNumber)

	if len(t.Name) == 0 {

		return errors.New("The user name does not a null string,try agin!")
	}
	if len(t.Nickname) == 0 {

		return errors.New("The nickname does not a null string,try agin!")
	}
	if len(t.Password) == 0 {

		return errors.New("The password does not a null string,try agin!")
	}

	exist, err := IsExistUserName(t.Name)

	if err != nil {
		return err
	}

	if exist {

		return errors.New("This user name is exist!")

	}
	exist, err = IsExistUserPhoneNumber(t.PhoneNumber)
	if err != nil {
		return err
	}
	if exist {

		return errors.New("This user PhoneNumber is exist!")

	}
	exist, err = IsExistUserEmail(t.Email)
	if err != nil {
		return err
	}
	if exist {

		return errors.New("This user Email is exist!")

	}
	return nil

}

//NewUsers 批量创建User
func NewUsers(ts []*User) ([]*User, error) {

	var err error
	tss := make([]interface{}, len(ts))

	for i, t := range ts {
		now := time.Now()
		bid := bson.NewObjectId()
		t.CreatedAt = &now
		t.ID = &bid
		t.UpdatedAt = &now
		t.DeletedAt = &time.Time{}
		t.Status = USERSTATUS_ACTIVATED
		if t.Type == USERTYPE_NULL {
			t.Type = USERTYPE_NORMAL
		}
		err := toCheckThisUser(t)
		if err != nil {
			return nil, err
		}
		password, salt := utils.GetFinalPwdWithSalt(t.Password)
		t.Password = password
		t.Salt = salt
		t.Token = bid.Hex() + utils.RandomString(16)
		tss[i] = t

	}

	err = CUser.Insert(tss...)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

//DeleteUsersByIDs 根据数据的id数组删除数据
func DeleteUsersByIDs(ids []*bson.ObjectId) bool {

	_, err := CUser.UpdateAll(bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}, bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	})

	return err == nil
}

//GetUserByIDs 根据数据的id数组查询数据
func GetUserByIDs(ids []*bson.ObjectId) ([]*User, error) {
	var ts []*User
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["_id"] = bson.M{
		"$in": ids,
	}
	err := CUser.Find(where).All(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//UpdateUsersByID 批量更新User
func UpdateUsersByID(ts []*User) ([]*User, error) {

	var err error
	tss := []*User{}

	for _, m := range ts {

		if !m.ID.Valid() {
			return nil, errors.New("The id is invalid,try agin!")
		}

		now := time.Now()
		m.UpdatedAt = &now

		tss = append(tss, m)
		update := make(bson.M)
		col := make(bson.M)

		update["$set"] = col

		if len(strings.TrimSpace(m.Name)) > 0 {

			col["name"] = m.Name
		}

		if m.Gender != GENDER_NULL {

			col["gender"] = m.Gender
		}

		if len(strings.TrimSpace(m.Nickname)) > 0 {

			col[`nickname`] = m.Nickname

		}
		if len(strings.TrimSpace(m.RealName)) > 0 {

			col[`real_name`] = m.Name

		}
		if len(strings.TrimSpace(m.Birthday)) > 0 {

			col[`birthday`] = m.Birthday
		}
		if len(strings.TrimSpace(m.Location)) > 0 {

			col[`location`] = m.Location
		}

		if len(strings.TrimSpace(m.PhoneNumber)) > 0 {

			col[`phone_number`] = m.PhoneNumber
		}
		if len(strings.TrimSpace(m.Password)) > 0 {
			password, salt := utils.GetFinalPwdWithSalt(m.Password)
			token := utils.RandomString(16)
			m.Token = token
			col[`password`] = password
			col[`salt`] = salt
			col[`token`] = m.ID.Hex() + token
		}

		if len(strings.TrimSpace(m.Email)) > 0 {

			col[`email`] = m.Email

		}

		if m.Type != USERTYPE_NULL {

			col[`type`] = m.Type

		}
		if m.Status != USERSTATUS_NULL {

			col[`status`] = m.Status

		}

		col[`deleted_at`] = &time.Time{}
		col[`updated_at`] = m.UpdatedAt

		err = CUser.UpdateId(m.ID, update)

		if err != nil {
			return nil, err
		}

	}

	return tss, nil
}

//GetUserWhereMap 根据byte数组获取map数据
func GetUserWhereMap(data []byte, where bson.M) bson.M {

	m := &User{}
	json.Unmarshal(data, m)

	if len(strings.TrimSpace(m.Name)) > 0 {

		where["name"] = m.Name
	}

	if m.Gender != GENDER_NULL {

		where["gender"] = m.Gender
	}

	if len(strings.TrimSpace(m.Nickname)) > 0 {

		where[`nickname`] = m.Nickname

	}
	if len(strings.TrimSpace(m.RealName)) > 0 {

		where[`real_name`] = m.Name

	}
	if len(strings.TrimSpace(m.Birthday)) > 0 {

		where[`birthday`] = m.Birthday
	}
	if len(strings.TrimSpace(m.Location)) > 0 {

		where[`location`] = m.Location
	}

	if len(strings.TrimSpace(m.PhoneNumber)) > 0 {

		where[`phone_number`] = m.PhoneNumber
	}

	if len(strings.TrimSpace(m.Email)) > 0 {

		where[`email`] = m.Email

	}

	if m.Type != USERTYPE_NULL {

		where[`type`] = m.Type

	}
	if m.Status != USERSTATUS_NULL {

		where[`status`] = m.Status

	}

	if m.ID != nil && m.ID.Valid() {

		where[`_id`] = m.ID

	}

	return where
}

//FindUsersByParams 根据查询参数获取User的分页数据
func FindUsersByParams(req *proto.RequestParams) ([]*User, int64, error) {

	var ts []*User

	offsetHead := req.OffsetHead
	OffsetTail := req.OffsetTail

	if OffsetTail-offsetHead <= 0 || OffsetTail-offsetHead > 200 {

		return nil, 0, errors.New("OffsetTail must be uper than offsetHead and  OffsetTail-offsetHead must lower 200!")
	}

	where := bson.M{}
	createT := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}

	if len(strings.TrimSpace(req.StartTime)) > 0 {
		st, _ := time.Parse(time.RFC3339, req.StartTime)
		createT["$gte"] = st
		where["created_at"] = createT

	}

	if len(strings.TrimSpace(req.EndTime)) > 0 {
		et, _ := time.Parse(time.RFC3339, req.EndTime)
		createT["$lte"] = et
		where["created_at"] = createT

	}

	if req.Where != nil && len(req.Where) > 0 {

		where = GetUserWhereMap(req.Where, where)
	}

	if len(strings.TrimSpace(req.Keywords)) > 0 {

		reg := bson.M{"$regex": bson.RegEx{req.Keywords, "i"}}
		or := []bson.M{bson.M{"name": reg}, bson.M{"nickname": reg}, bson.M{"real_name": reg}, bson.M{"phone_number": reg}, bson.M{"email": reg}}
		where["$or"] = or

	}

	query := CUser.Find(where)

	if len(strings.TrimSpace(req.OrderBy)) > 0 {
		if req.IsDesc {
			query = query.Sort("-" + req.OrderBy)
		} else {
			query = query.Sort(req.OrderBy)
		}
	} else {
		query = query.Sort("-created_at")
	}

	err := query.Skip(int(offsetHead)).Limit(int(OffsetTail - offsetHead)).All(&ts)

	if err != nil {

		return nil, 0, err
	}

	//totalNum mongodb中所有User的总数，而非查询返回的条数
	totalNum, _ := CUser.Find(where).Count()

	return ts, int64(totalNum), nil
}

//GetUserByName 根据UserName查询数据
func GetUserByName(name string) (u *User, err error) {
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["name"] = bson.RegEx{name, "i"}
	err = CUser.Find(where).One(&u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

//GetUserByID 根据UserID查询数据
func GetUserByID(id bson.ObjectId) (u *User, err error) {
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["_id"] = id
	err = CUser.Find(where).One(&u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

//GetUserByToken 根据UserToken查询数据
func GetUserByToken(token string) (u *User, err error) {
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["token"] = token
	err = CUser.Find(where).One(&u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

//GetUserByPhoneNumber 根据UserPhoneNumber查询数据
func GetUserByPhoneNumber(phoneNum string) (u *User, err error) {
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["phone_number"] = phoneNum
	err = CUser.Find(where).One(&u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

//GetUserByEmail 根据UserEmail查询数据
func GetUserByEmail(email string) (u *User, err error) {
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["email"] = bson.RegEx{email, "i"}
	err = CUser.Find(where).One(&u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

//IsExistUserName 根据UserName查询数据
func IsExistUserName(name string) (bool, error) {

	if len(name) <= 4 || len(name) > 10 {

		return false, errors.New("This length of user name  must be with 5 - 10 !")
	}
	u, err := GetUserByName(name)
	if err != nil {
		return u != nil, err
	}
	return u != nil, nil
}

//IsExistUserPhoneNumber 根据UserPhoneNumber查询数据
func IsExistUserPhoneNumber(phoneNum string) (bool, error) {
	phoneNum = strings.TrimSpace(phoneNum)
	if len(phoneNum) == 0 {

		return false, nil
	}

	if !utils.IsMobile(phoneNum) {
		return false, errors.New("This mobile field is invalid!")
	}

	u, err := GetUserByPhoneNumber(phoneNum)
	if err != nil {
		return u != nil, err
	}
	return u != nil, nil
}

//IsExistUserEmail 根据UserEmail查询数据
func IsExistUserEmail(email string) (bool, error) {
	email = strings.TrimSpace(email)
	if len(email) == 0 {

		return false, nil
	}
	if !utils.IsEmail(email) {
		return false, errors.New("This email field is invalid!")
	}
	u, err := GetUserByEmail(email)
	if err != nil {
		return u != nil, err
	}
	return u != nil, nil
}

//IsExistUserEmail 根据关键信息查询数据用户数据
func IsExistUser(name string) (bool, error) {

	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}

	or := []bson.M{bson.M{"name": bson.RegEx{name, "i"}}, bson.M{"phone_number": bson.RegEx{name, "i"}}, bson.M{"email": bson.RegEx{name, "i"}}}
	where["$or"] = or

	u := &User{}
	err := CUser.Find(where).One(u)
	if err != nil {
		if err == mgo.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return u != nil, nil
}
