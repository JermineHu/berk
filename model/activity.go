package model

import (
	"errors"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"time"
	"reflect"
	"gopkg.in/mgo.v2"
	"git.vdo.space/foy/config"
	"git.vdo.space/foy/model"
)

type Activity struct {
	ID          bson.ObjectId        `bson:"_id" json:"id"`
	GroupID     string               `bson:"group_id" json:"group_id"`
	ImgName     string               `json:"img_name" bson:"img_name"`
	ToURL       string               `json:"to_url" bson:"to_url"`
	Title       string               `json:"title,omitempty" bson:"title,omitempty"`
	Description string               `json:"description,omitempty" bson:"description,omitempty"`
	JoinsCount  int64                `json:"joins_count,omitempty" bson:"joins_count,omitempty"`
	StartTime   time.Time            `json:"start_time,omitempty" bson:"start_time,omitempty"`
	EndTime     time.Time            `json:"end_time,omitempty" bson:"end_time,omitempty"`
	Status      proto.ActivityStatus `bson:"status,omitempty" json:"status,omitempty"`
	Type        proto.ActivityType   `bson:"type,omitempty" json:"type,omitempty"`
	Region      []string             `bson:"region,omitempty" json:"region,omitempty"`
	Lang        string               `bson:"lang,omitempty" json:"lang,omitempty"`
	CreatedAt   *time.Time           `bson:"created_at" json:"created_at"`
	UpdatedAt   *time.Time           `bson:"updated_at" json:"updated_at"`
	DeletedAt   *time.Time           `bson:"deleted_at" json:"deleted_at"`
}

func (v Activity) GetCollectionName() string {

	return reflect.TypeOf(v).Name()
}

//GetActivityStatusByStr 根据字符串获取数据
func GetActivityStatusByStr(str string) proto.ActivityStatus {

	switch str {

	case proto.ActivityStatus_ActivityStatusOnline.String():

		return proto.ActivityStatus_ActivityStatusOnline

		break
	case proto.ActivityStatus_ActivityStatusOffline.String():

		return proto.ActivityStatus_ActivityStatusOffline

		break

	}

	return proto.ActivityStatus_ActivityStatusNULL

}

//GetActivityTypeByStr 根据字符串获取数据
func GetActivityTypeByStr(str string) proto.ActivityType {

	switch str {

	case proto.ActivityType_ActivityTypeFilter.String():

		return proto.ActivityType_ActivityTypeFilter

		break
	case proto.ActivityType_ActivityTypeOthers.String():

		return proto.ActivityType_ActivityTypeOthers

		break

	}

	return proto.ActivityType_ActivityTypeNULL

}

// VaildDataImageFilter 验证数据
func VaildDataActivity(t *Activity) error {

	if len(strings.TrimSpace(t.Title)) == 0 {

		return errors.New("The Title does not a null string,try agin!")
	}
	if len(strings.TrimSpace(t.ImgName)) == 0 {

		return errors.New("The ImgName does not a null string,try agin!")
	}

	if t.StartTime.IsZero() {

		return errors.New("The StartTime does not a null string,try agin!")
	}

	if t.EndTime.IsZero() {

		return errors.New("The EndTime does not a null string,try agin!")
	}
	if len(strings.TrimSpace(t.Lang)) == 0 {

		return errors.New("The Language does not a null string,try agin!")
	}

	return nil

}

func ActivityToData(md *Activity) (model.Dungeons, error) {
	d := make(model.Dungeons)
	d[`id`] = md.ID
	if len(strings.TrimSpace(md.ImgName)) > 0 {
		d[`image_url`] = config.GetResourceAddr() + "activities/" + md.ImgName
	}
	d[`to_url`] = md.ToURL
	d[`img_name`] = md.ImgName
	d[`title`] = md.Title
	d[`lang`] = md.Lang
	d[`region`] = md.Region
	d[`description`] = md.Description
	d[`joins_count`] = md.JoinsCount
	d[`start_time`] = md.StartTime.In(time.Local).Format(time.RFC3339)
	d[`end_time`] = md.EndTime.In(time.Local).Format(time.RFC3339)
	d[`status`] = md.Status.String()
	d[`type`] = md.Type.String()
	d[`created_at`] = md.CreatedAt
	d[`updated_at`] = md.UpdatedAt
	d[`display_duration`] = 4

	return d, nil
}

func ActivitysToDatas(mds []*Activity) ([]model.Dungeons, error) {
	dl := []model.Dungeons{}
	for _, x := range mds {
		d, err := ActivityToData(x)
		if err != nil {
			return nil, err
		}
		dl = append(dl, d)
	}

	return dl, nil
}


//NewActivitys 批量创建Activity
func NewActivitys(ts []*Activity) ([]*Activity, error) {

	var err error
	tss := make([]interface{}, len(ts))
	groupID := bson.NewObjectId().Hex()

	for i, t := range ts {
		now := time.Now()
		t.CreatedAt = &now
		t.ID = bson.NewObjectId()
		t.UpdatedAt = &now
		t.DeletedAt = &time.Time{}
		//t.Status = proto.ActivityStatus_ActivityStatusOffline
		if len(strings.TrimSpace(t.GroupID)) == 0 {
			t.GroupID = groupID
		}
		err := VaildDataActivity(t)
		if err != nil {
			return nil, err
		}
		tss[i] = t
	}

	err = CActivity.Insert(tss...)

	if err != nil {
		return nil, err
	}

	return ts, nil
}

//DeleteActivitysByIDs 根据数据的id数组删除数据
func DeleteActivitysByIDs(ids []*bson.ObjectId) bool {

	_, err := CActivity.UpdateAll(bson.M{
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


//GetActivityByID 根据数据的id查询数据
func GetActivityByID(id *bson.ObjectId) (*Activity, error) {
	var ts *Activity
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["_id"] = id

	err := CActivity.Find(where).One(&ts)
	if err != nil && err != mgo.ErrNotFound {

		return nil, err
	}
	return ts, nil
}

//GetActivityByIDs 根据数据的id数组查询数据
func GetActivityByIDs(ids []*bson.ObjectId) ([]*Activity, error) {
	var ts []*Activity
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["_id"] = bson.M{
		"$in": ids,
	}
	err := CActivity.Find(where).All(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//UpdateActivitysByID 批量更新Activity
func UpdateActivitysByID(ts []*Activity) ([]*Activity, error) {

	var err error
	tss := []*Activity{}

	for _, m := range ts {

		now := time.Now()
		m.UpdatedAt = &now

		tss = append(tss, m)
		update := make(bson.M)
		col := make(bson.M)

		update["$set"] = col

		if len(strings.TrimSpace(m.ToURL)) > 0 {

			col["to_url"] = m.ToURL
		}

		if len(strings.TrimSpace(m.ImgName)) > 0 {

			col["img_name"] = m.ImgName
		}

		if len(strings.TrimSpace(m.Title)) > 0 {

			col[`title`] = m.Title

		}
		if m.Status != proto.ActivityStatus_ActivityStatusNULL {

			col[`status`] = m.Status

		}
		if len(strings.TrimSpace(m.Description)) > 0 {

			col[`description`] = m.Description
		}

		if m.Type != proto.ActivityType_ActivityTypeNULL {

			col[`type`] = m.Type
		}
		if !m.EndTime.IsZero() {

			col[`end_time`] = m.EndTime
		}

		if !m.StartTime.IsZero() {

			col[`start_time`] = m.StartTime

		}
		if len(strings.TrimSpace(m.Lang)) > 0 {

			col["lang"] = m.Lang
		}
		if len(strings.TrimSpace(m.GroupID)) > 0 {

			col["group_id"] = m.GroupID
		}
		if len(m.Region) > 0 {

			col["region"] = m.Region
		}

		col[`deleted_at`] = &time.Time{}
		col[`updated_at`] = m.UpdatedAt

		err = CActivity.UpdateId(m.ID, update)

		if err != nil {
			return nil, err
		}

	}

	return tss, nil
}



//FindActivitysByParams 根据查询参数获取Activity的分页数据
func FindActivitysByParams(req *proto.RequestParams) ([]*Activity, int64, error) {

	var ts []*Activity

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

		//where = GetActivityWhereMap(req.Where, where)
	}

	if len(strings.TrimSpace(req.Keywords)) > 0 {

		reg := bson.M{"$regex": bson.RegEx{req.Keywords, "i"}}
		or := []bson.M{bson.M{"title": reg}, bson.M{"description": reg}}
		where["$or"] = or

	}

	query := CActivity.Find(where)

	if len(strings.TrimSpace(req.OrderBy)) > 0 {
		if req.IsDesc {
			query = query.Sort("-" + req.OrderBy)
		} else {
			query = query.Sort(req.OrderBy)
		}
	} else {
		//query = query.Sort("-created_at")
		query = query.Sort("-start_time")
	}

	err := query.Skip(int(offsetHead)).Limit(int(OffsetTail - offsetHead)).All(&ts)

	if err != nil {

		return nil, 0, err
	}

	//totalNum mongodb中所有Activity的总数，而非查询返回的条数
	totalNum, _ := CActivity.Find(where).Count()

	return ts, int64(totalNum), nil
}
