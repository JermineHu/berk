package model

import (
	"errors"
	"git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
	"time"
	"encoding/json"
)

type TagRequestQueryParams struct {
	model.RequestQueryParams
	Where Tag `json:"where"`
}

type Tag struct {
	ID          bson.ObjectId  `bson:"_id" json:"id"`
	UserID      *bson.ObjectId `bson:"user_id" json:"user_id"`
	Name        string         `bson:"name" json:"name"`
	Color       string         `bson:"color" json:"color"`
	Description string         `bson:"description" json:"description"`
	CreatedAt   *time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   *time.Time     `bson:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time     `bson:"deleted_at" json:"deleted_at"`
}

func (v Tag) GetCollectionName() string {
	return strings.ToLower(reflect.TypeOf(v).Name())
}
func TagToData(md *Tag) model.Dungeons {
	d := make(model.Dungeons)
	d[`id`] = md.ID
	d[`name`] = md.Name
	d[`color`] = md.Color
	return d
}

func TagsToDatas(mds []*Tag) []model.Dungeons {
	dl := []model.Dungeons{}
	for _, x := range mds {
		dl = append(dl, TagToData(x))
	}
	return dl
}

//IsExistTag 根据关键信息查询数据
func IsExistTag(name  string, userID *bson.ObjectId) (bool, error) {

	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}

	where["name"] = name
	where["user_id"] = userID

	u := &Tag{}
	err := CTag.Find(where).One(u)

	if err != nil {
		if err == mgo.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return u != nil, nil
}

//NewTags 批量创建标签
func NewTags(ts []*Tag) ([]interface{}, error) {

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

		exist, _ := IsExistTag(t.Name, t.UserID)

		if exist {

			return nil, errors.New("The name already exist,try agin!")

		}

		tss[i] = t
	}

	err = CTag.Insert(tss...)

	if err != nil {
		return nil, err
	}

	return tss, nil
}

//DeleteByIDs 根据数据的id数组删除数据
func DeleteTagByIDs(ids []*bson.ObjectId, userID *bson.ObjectId) bool {

	_, err := CTag.UpdateAll(bson.M{
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

//GetTagByIDs 根据数据的id数组查询数据
func GetTagByIDs(ids []*bson.ObjectId, userID *bson.ObjectId) ([]*Tag, error) {
	var ts []*Tag
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

	err := CTag.Find(where).All(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//UpdateTags 批量更新标签
func UpdateTags(ts []*Tag) ([]*Tag, error) {

	var err error
	tss := []*Tag{}

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
		if len(strings.TrimSpace(t.Color)) > 0 {

			col["color"] = t.Color
		}

		err = CTag.Update(bson.D{{"_id", t.ID}, {"user_id", t.UserID}}, update)

	}

	if err != nil {
		return nil, err
	}

	return tss, nil
}

//GetTagWhereMap 根据byte数组获取map数据
func GetTagWhereMap(data []byte, where bson.M) bson.M {

	m := &Tag{}
	json.Unmarshal(data, m)

	if len(strings.TrimSpace(m.Name)) > 0 {

		where["name"] = m.Name
	}
	if len(strings.TrimSpace(m.Color)) > 0 {

		where["color"] = m.Color
	}

	if m.ID.Valid() {

		where[`_id`] = m.ID

	}

	if m.UserID != nil && m.UserID.Valid() {

		where[`user_id`] = m.UserID

	}

	return where
}

//FindTagsByParams 根据查询参数获取标签的分页数据
func FindTagsByParams(req *proto.RequestParams, userID *bson.ObjectId) ([]*Tag, int64, error) {

	var ts []*Tag

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

	if len(req.Where) > 0 {

		where = GetTagWhereMap(req.Where, where)
	}

	if len(strings.TrimSpace(req.Keywords)) > 0 {
		reg := bson.M{"$regex": bson.RegEx{req.Keywords, "i"}}
		or := []bson.M{bson.M{"name": reg}, bson.M{"color": reg}, bson.M{"description": reg}}
		where["$or"] = or

	}

	where[`user_id`] = userID
	query := CTag.Find(where)

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
	totalNum, _ := CTag.Find(where).Count()

	return ts, int64(totalNum), nil
}
