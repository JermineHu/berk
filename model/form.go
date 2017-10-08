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

type Form struct {
	ID                 bson.ObjectId  `bson:"_id" json:"id"`
	UserID             *bson.ObjectId `bson:"user_id" json:"user_id"`
	VistorName         string         `bson:"vistor_name" json:"vistor_name" binding:"required"`
	CompanyName        string         `bson:"company_name" json:"company_name" binding:"required"`
	CompanyAddress     string         `bson:"company_address" json:"company_address" binding:"required"`
	CompanyContactType string         `bson:"company_contact_type" json:"company_contact_type" binding:"required"`
	CompanyContactInfo string         `bson:"company_contact_info" json:"company_contact_info" binding:"required"`
	Description        string         `bson:"description" json:"description" binding:"required"`
	CreatedAt          *time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt          *time.Time     `bson:"updated_at" json:"updated_at"`
	DeletedAt          *time.Time     `bson:"deleted_at" json:"deleted_at"`
}

func (v Form) GetCollectionName() string {
	return strings.ToLower(reflect.TypeOf(v).Name())
}
func FormToData(md *Form) model.Dungeons {
	d := make(model.Dungeons)
	d[`id`] = md.ID
	d[`name`] = md.CompanyName
	return d
}

func FormsToDatas(mds []*Form) []model.Dungeons {
	dl := []model.Dungeons{}
	for _, x := range mds {
		dl = append(dl, FormToData(x))
	}
	return dl
}

//NewForms 批量创建标签
func NewForms(ts []*Form) ([]interface{}, error) {

	var err error
	tss := make([]interface{}, len(ts))

	for i, t := range ts {
		now := time.Now()
		t.CreatedAt = &now
		t.ID = bson.NewObjectId()
		t.UpdatedAt = &now
		t.DeletedAt = &time.Time{}

		if len(strings.TrimSpace(t.CompanyName)) == 0 {

			return nil, errors.New("The name does not a null string,try agin!")
		}

		tss[i] = t
	}

	err = CForm.Insert(tss...)

	if err != nil {
		return nil, err
	}

	return tss, nil
}

//DeleteByIDs 根据数据的id数组删除数据
func DeleteFormByIDs(ids []*bson.ObjectId, userID *bson.ObjectId) bool {

	_, err := CForm.UpdateAll(bson.M{
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

//GetFormByIDs 根据数据的id数组查询数据
func GetFormByIDs(ids []*bson.ObjectId, userID *bson.ObjectId) ([]*Form, error) {
	var ts []*Form
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

	err := CForm.Find(where).All(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//UpdateForms 批量更新
func UpdateForms(ts []*Form) ([]*Form, error) {

	var err error
	tss := []*Form{}

	for _, t := range ts {

		if !t.ID.Valid() {

			return nil, errors.New("The id does not a null string,try agin!")
		}

		if len(strings.TrimSpace(t.CompanyName)) == 0 {

			return nil, errors.New("The name does not a null string,try agin!")
		}
		now := time.Now()
		t.UpdatedAt = &now

		tss = append(tss, t)
		update := make(bson.M)
		col := make(bson.M)

		update["$set"] = col

		col["name"] = t.CompanyName

		if len(strings.TrimSpace(t.Description)) > 0 {

			col["description"] = t.Description
		}

		err = CForm.Update(bson.D{{"_id", t.ID}, {"user_id", t.UserID}}, update)

	}

	if err != nil {
		return nil, err
	}

	return tss, nil
}

//FindFormsByParams 根据查询参数获取标签的分页数据
func FindFormsByParams(req *proto.RequestParams, userID *bson.ObjectId) ([]*Form, int64, error) {

	var ts []*Form

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
	query := CForm.Find(where)

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
	totalNum, _ := CForm.Find(where).Count()

	return ts, int64(totalNum), nil
}
