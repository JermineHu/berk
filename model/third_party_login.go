package model

import (
	"errors"
	"git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
	"time"
	"git.vdo.space/foy/utils"
	"gopkg.in/mgo.v2"
)

type LoginType int

const (
	LoginType_NULL LoginType = 0
	LoginType_WECHAT LoginType = 1
	LoginType_QQ LoginType = 2
)

type ThirdPartyLogin struct {
	ID             bson.ObjectId  `bson:"_id" json:"id"`
	UserID         *bson.ObjectId `bson:"user_id" json:"user_id"`
	UnionID        string         `bson:"union_id" json:"union_id"`
	LotteryDrawNum string         `bson:"lottery_draw_num" json:"lottery_draw_num"`
	LoginType      LoginType         `bson:"login_type" json:"login_type"`
	CreatedAt      *time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt      *time.Time     `bson:"updated_at" json:"updated_at"`
	DeletedAt      *time.Time     `bson:"deleted_at" json:"deleted_at"`
}

func (v ThirdPartyLogin) GetCollectionName() string {
	return strings.ToLower(reflect.TypeOf(v).Name())
}

func ThirdPartyLoginToData(md *ThirdPartyLogin) (model.Dungeons, error) {
	d := make(model.Dungeons)
	d[`id`] = md.ID
	d[`union_id`] = md.UnionID
	d[`login_type`] = md.LoginType
	return d, nil
}

func ThirdPartyLoginsToDatas(mds []*ThirdPartyLogin) ([]model.Dungeons, error) {
	dl := []model.Dungeons{}
	for _, x := range mds {
		d, err := ThirdPartyLoginToData(x)
		if err != nil {
			return nil, err
		}
		dl = append(dl, d)
	}

	return dl, nil
}

//NewThirdPartyLogins 批量创建标签
func NewThirdPartyLogins(ts []*ThirdPartyLogin) ([]interface{}, error) {

	var err error
	tss := make([]interface{}, len(ts))

	for i, t := range ts {
		now := time.Now()
		t.CreatedAt = &now
		t.ID = bson.NewObjectId()
		t.UpdatedAt = &now
		t.DeletedAt = &time.Time{}
		tss[i] = t
	}

	err = CThirdPartyLogin.Insert(tss...)

	if err != nil {
		return nil, err
	}

	return tss, nil
}

//DeleteByIDs 根据数据的id数组删除数据
func DeleteThirdPartyLoginByIDs(ids []*bson.ObjectId, userID *bson.ObjectId) bool {

	_, err := CThirdPartyLogin.UpdateAll(bson.M{
		"_id": bson.M{
			"$in": ids,
		},
		"user_id": userID,
	}, bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	})

	return err == nil
}

//GetThirdPartyLoginByIDs 根据数据的id数组查询数据
func GetThirdPartyLoginByIDs(ids []*bson.ObjectId, userID *bson.ObjectId) ([]*ThirdPartyLogin, error) {
	var ts []*ThirdPartyLogin
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["_id"] = bson.M{
		"$in": ids,
	}
	if userID != nil {
		where[`user_id`] = userID
	}

	err := CThirdPartyLogin.Find(where).All(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//GetThirdPartyLoginByIDs 根据数据的id数组查询数据
func GetThirdPartyLoginByUserIDAndType(userID *bson.ObjectId,lt LoginType) (*ThirdPartyLogin, error) {
	var ts *ThirdPartyLogin
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	if userID != nil {
		where[`user_id`] = userID
	}
	if lt != LoginType_NULL {
		where[`login_type`] = lt
	}

	err := CThirdPartyLogin.Find(where).One(&ts)
	if err != nil {

		if err==mgo.ErrNotFound {
			return nil,nil
		}

		return nil, err
	}
	return ts, nil
}

//FindThirdPartyLoginsByParams 根据查询参数获取标签的分页数据
func FindThirdPartyLoginsByParams(req *proto.RequestParams, lt LoginType) ([]*ThirdPartyLogin, int64, error) {

	var ts []*ThirdPartyLogin

	offsetHead := req.OffsetHead
	OffsetTail := req.OffsetTail

	if OffsetTail - offsetHead <= 0 || OffsetTail - offsetHead > 200 {

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

	if len(req.UserUUID) > 0 {
		where[`user_id`] = utils.Str2ObjectId(req.UserUUID)
	}

	if lt != LoginType_NULL {
		where[`login_type`] = lt
	}

	query := CThirdPartyLogin.Find(where)

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

	//totalNum mongodb中所有标签的总数，而非查询返回的条数
	totalNum, _ := CThirdPartyLogin.Find(where).Count()

	return ts, int64(totalNum), nil
}
