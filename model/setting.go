package model

import (
	"errors"
	"git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
	"time"
)

type Setting struct {
	ID          bson.ObjectId  `bson:"_id" json:"id"`
	UserID      *bson.ObjectId `bson:"user_id" json:"user_id"`
	Name        string         `bson:"name" json:"name"`
	Description string         `bson:"description" json:"description"`
	CreatedAt   *time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   *time.Time     `bson:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time     `bson:"deleted_at" json:"deleted_at"`
}

func (v Setting) GetCollectionName() string {
	return strings.ToLower(reflect.TypeOf(v).Name())
}
func SettingToData(md *Setting) (model.Dungeons, error) {
	d := make(model.Dungeons)
	d[`id`] = md.ID
	d[`name`] = md.Name
	return d, nil
}

func SettingsToDatas(mds []*Setting) ([]model.Dungeons, error) {
	dl := []model.Dungeons{}
	for _, x := range mds {
		d, err := SettingToData(x)
		if err != nil {
			return nil, err
		}
		dl = append(dl, d)
	}

	return dl, nil
}

//NewSettings 批量创建标签
func NewSettings(ts []*Setting) ([]interface{}, error) {

	var err error
	tss := make([]interface{}, len(ts))

	for i, t := range ts {
		now := time.Now()
		t.CreatedAt = &now
		t.ID = bson.NewObjectId()
		t.UpdatedAt = &now
		t.DeletedAt = &time.Time{}

		if len(strings.TrimSpace(t.Name)) == 0 {

			return nil, errors.New("The name does not a null string,try agin!")
		}

		tss[i] = t
	}

	err = CSetting.Insert(tss...)

	if err != nil {
		return nil, err
	}

	return tss, nil
}

//DeleteByIDs 根据数据的id数组删除数据
func DeleteSettingByIDs(ids []*bson.ObjectId, userID *bson.ObjectId) bool {

	_, err := CSetting.UpdateAll(bson.M{
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

//GetSettingByIDs 根据数据的id数组查询数据
func GetSettingByIDs(ids []*bson.ObjectId, userID *bson.ObjectId) ([]*Setting, error) {
	var ts []*Setting
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

	err := CSetting.Find(where).All(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//UpdateSettings 批量更新标签
func UpdateSettings(ts []*Setting) ([]*Setting, error) {

	var err error
	tss := []*Setting{}

	for _, t := range ts {

		if !t.ID.Valid() {

			return nil, errors.New("The id does not a null string,try agin!")
		}

		if len(strings.TrimSpace(t.Name)) == 0 {

			return nil, errors.New("The name does not a null string,try agin!")
		}
		now := time.Now()
		t.UpdatedAt = &now

		tss = append(tss, t)
		update := make(bson.M)
		col := make(bson.M)

		update["$set"] = col

		col["name"] = t.Name

		if len(strings.TrimSpace(t.Description)) > 0 {

			col["description"] = t.Description
		}

		err = CSetting.Update(bson.D{{"_id", t.ID}, {"user_id", t.UserID}}, update)

	}

	if err != nil {
		return nil, err
	}

	return tss, nil
}

//FindSettingsByParams 根据查询参数获取标签的分页数据
func FindSettingsByParams(req *proto.RequestParams, userID *bson.ObjectId) ([]*Setting, int64, error) {

	var ts []*Setting

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
	where[`user_id`] = userID
	query := CSetting.Find(where)

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
	totalNum, _ := CSetting.Find(where).Count()

	return ts, int64(totalNum), nil
}
