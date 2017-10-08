package model

import (
	"encoding/json"
	"errors"
	"git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
	"time"
)

type Feedback struct {
	ID          bson.ObjectId      `bson:"_id" json:"id"`
	VistorID    string             `bson:"vistor_id" json:"vistor_id" `
	VistorName  string             `bson:"vistor_name" json:"vistor_name"`
	Description string             `bson:"description" json:"description"`
	ExtraInfo   interface{}        `bson:"extra_info" json:"extra_info"`
	CreatedAt   *time.Time         `bson:"created_at" json:"created_at"`
	UpdatedAt   *time.Time         `bson:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time         `bson:"deleted_at" json:"deleted_at"`
	IsShow      *proto.BooleanType `bson:"is_show" json:"is_show"`
}

func (v Feedback) GetCollectionName() string {
	return strings.ToLower(reflect.TypeOf(v).Name())
}
func FeedbackToData(md *Feedback) model.Dungeons {
	d := make(model.Dungeons)
	d[`id`] = md.ID
	d[`vistor_name`] = md.VistorName
	d[`description`] = md.Description
	d[`extra_info`] = md.ExtraInfo
	return d
}

func FeedbacksToDatas(mds []*Feedback) []model.Dungeons {
	dl := []model.Dungeons{}
	for _, x := range mds {
		dl = append(dl, FeedbackToData(x))
	}
	return dl
}

//NewFeedbacks 批量创建标签
func NewFeedbacks(ts []*Feedback) ([]interface{}, error) {

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

	err = CFeedback.Insert(tss...)

	if err != nil {
		return nil, err
	}

	return tss, nil
}

//DeleteByIDs 根据数据的id数组删除数据
func DeleteFeedbackByIDs(ids []*bson.ObjectId) bool {

	_, err := CFeedback.UpdateAll(bson.M{
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

//GetFeedbackByIDs 根据数据的id数组查询数据
func GetFeedbackByIDs(ids []*bson.ObjectId) ([]*Feedback, error) {
	var ts []*Feedback
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["_id"] = bson.M{
		"$in": ids,
	}
	err := CFeedback.Find(where).All(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//GetFeedbackByOpenID 根据数据的openid查询数据
func GetFeedbackByOpenID(openid string) ([]*Feedback, error) {
	var ts []*Feedback
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["vistor_id"] = openid

	err := CFeedback.Find(where).All(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//UpdateFeedbacks 批量更新
func UpdateFeedbacks(ts []*Feedback) ([]*Feedback, error) {

	var err error
	tss := []*Feedback{}

	for _, t := range ts {

		if !t.ID.Valid() {

			return nil, errors.New("The id does not a null string,try agin!")
		}
		now := time.Now()
		t.UpdatedAt = &now

		tss = append(tss, t)
		update := make(bson.M)
		col := make(bson.M)

		update["$set"] = col

		if len(strings.TrimSpace(t.Description)) > 0 {

			col["description"] = t.Description
		}
		if len(strings.TrimSpace(t.VistorName)) > 0 {

			col["vistor_name"] = t.VistorName
		}
		if t.ExtraInfo != nil {

			col["extra_info"] = t.ExtraInfo
		}

		if t.IsShow != nil && !t.IsShow.Null {

			col["is_show"] = t.IsShow
		}

		err = CFeedback.Update(bson.D{{"_id", t.ID}}, update)

	}

	if err != nil {
		return nil, err
	}

	return tss, nil
}

//GetFeedbackWhereMap 根据byte数组获取map数据
func GetFeedbackWhereMap(data []byte, where bson.M) bson.M {

	m := &Feedback{}
	json.Unmarshal(data, m)

	if m.IsShow != nil {

		where["is_show.bool"] = m.IsShow.Bool
	}

	return where
}

//FindFeedbacksByParams 根据查询参数获取标签的分页数据
func FindFeedbacksByParams(req *proto.RequestParams, isFront bool) ([]*Feedback, int64, error) {

	var ts []*Feedback

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
		where = GetFeedbackWhereMap(req.Where, where)
	}

	if isFront {
		// is frontend
		where["is_show.bool"] = isFront
	}

	query := CFeedback.Find(where)

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
	totalNum, _ := CFeedback.Find(where).Count()

	return ts, int64(totalNum), nil
}
