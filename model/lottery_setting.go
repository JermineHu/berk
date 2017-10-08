package model

import (
	"encoding/json"
	"errors"
	. "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
	"time"
	"gopkg.in/mgo.v2"
	"math/rand"
	"git.vdo.space/foy/config"
)

type LotteryType int32

const (
	LotteryTypeNULL LotteryType = 0
	LotteryTypeVDOIndoor LotteryType = 1
	LotteryTypeVDOOutdoor LotteryType = 2
	LotteryTypeVDO20170710 LotteryType = 3
)

//GetStrByLotteryType 获取类型的string
func GetStrByLotteryType(status LotteryType) string {

	switch status {

	case LotteryTypeVDOIndoor:

		return "LotteryTypeVDOIndoor"

		break

	case LotteryTypeVDOOutdoor:

		return "LotteryTypeVDOOutdoor"
	case LotteryTypeVDO20170710:

		return "LotteryTypeVDO20170710"

		break
	}

	return "LotteryTypeNULL"

}

//GetLotteryTypeByStr 根据string获取类型
func GetLotteryTypeByStr(status string) LotteryType {

	switch status {

	case "LotteryTypeVDOIndoor":

		return LotteryTypeVDOIndoor

		break

	case "LotteryTypeVDOOutdoor":

		return LotteryTypeVDOOutdoor
	case "LotteryTypeVDO20170710":

		return LotteryTypeVDO20170710

		break
	}

	return LotteryTypeNULL

}

type LotteryLevel int32

const (
	LotteryLevelNULL LotteryLevel = 0
	LotteryLevelFirst LotteryLevel = 1
	LotteryLevelSecond LotteryLevel = 2
)

//GetStrByLotteryLevel 获取状态的string
func GetStrByLotteryLevel(status LotteryLevel) string {

	switch status {

	case LotteryLevelFirst:

		return "LotteryLevelFirst"

		break

	case LotteryLevelSecond:

		return "LotteryLevelSecond"

		break
	}

	return "LotteryLevelNULL"

}

//GetLotteryLevelByStr 根据string获取状态
func GetLotteryLevelByStr(status string) LotteryLevel {

	switch status {

	case "LotteryLevelFirst":

		return LotteryLevelFirst

		break

	case "LotteryLevelSecond":

		return LotteryLevelSecond

		break
	}

	return LotteryLevelNULL

}

type LotterySetting struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`                                        //编号
	LotteryName  string        `json:"lottery_name,omitempty" bson:"lottery_name,omitempty"` //奖品名称
	GiftCount    int64         `json:"gift_count,omitempty" bson:"gift_count,omitempty"`     // 奖品数量
	StartTime    int64         `json:"start_time,omitempty" bson:"start_time,omitempty"`     // 开始时间
	EndTime      int64         `json:"end_time,omitempty" bson:"end_time,omitempty"`         //结束时间
	LotteryExp   int64     `json:"lottery_exp,omitempty" bson:"lottery_exp,omitempty"`       //有效截止日期
	Chance       float64       `json:"chance" bson:"chance"`                                 //中奖概率
	LotteryType  LotteryType   `bson:"lottery_type,omitempty" json:"lottery_type,omitempty"` // 抽奖类型
	LotteryLevel LotteryLevel  `bson:"lottery_level,omitempty" json:"lottery_level"`         // 抽奖级别
	CreatedAt    *time.Time    `bson:"created_at" json:"created_at"`
	UpdatedAt    *time.Time    `bson:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time    `bson:"deleted_at" json:"deleted_at"`
}

func (v LotterySetting) GetCollectionName() string {

	return reflect.TypeOf(v).Name()
}

func LotterySettingToData(md *LotterySetting) (Dungeons, error) {
	d := make(Dungeons)
	d[`id`] = md.ID
	d[`lottery_name`] = md.LotteryName
	d[`gift_count`] = md.GiftCount
	d[`start_time`] = md.StartTime
	d[`end_time`] = md.EndTime
	d[`lottery_exp`] = md.LotteryExp
	d[`chance`] = md.Chance
	d[`lottery_type`] = GetStrByLotteryType(md.LotteryType)
	d[`lottery_level`] = GetStrByLotteryLevel(md.LotteryLevel)

	return d, nil
}

func LotterySettingsToDatas(mds []*LotterySetting) ([]Dungeons, error) {
	dl := []Dungeons{}
	for _, x := range mds {
		d, err := LotterySettingToData(x)
		if err != nil {
			return nil, err
		}
		dl = append(dl, d)
	}

	return dl, nil
}

// VaildDataImageFilter 验证数据
func VaildDataLotterySetting(t *LotterySetting) error {

	if len(strings.TrimSpace(t.LotteryName)) == 0 {

		return errors.New("The LotteryName does not a null string,try agin!")
	}
	if !t.ID.Valid() {

		return errors.New("The ID was needed,try agin!")
	}
	if t.GiftCount == 0 {

		return errors.New("The GiftCount not with 0,try agin!")
	}
	if t.StartTime == 0 {

		return errors.New("The StartTime was needed,try agin!")
	}
	if t.EndTime == 0 {

		return errors.New("The EndTime was needed,try agin!")
	}
	if t.LotteryExp == 0 {

		return errors.New("The LotteryExp was needed,try agin!")
	}
	if t.GiftCount == 0 {

		return errors.New("The GiftCount not with 0,try agin!")
	}
	if t.Chance == float64(0) {

		return errors.New("The Rate not with 0,try agin!")
	}

	return nil

}

type LotterySettingNum struct {
	ID               int64 `bson:"_id" json:"id"`
	LotterySettingID bson.ObjectId `bson:"lottery_setting_id" json:"lottery_setting_id"`
	LotteryLevel     LotteryLevel  `bson:"lottery_level,omitempty" json:"lottery_level"` // 抽奖级别
	LotteryType  LotteryType   `bson:"lottery_type,omitempty" json:"lottery_type,omitempty"` // 抽奖类型
}

func (v LotterySettingNum) GetCollectionName() string {

	return strings.ToLower(reflect.TypeOf(v).Name())
}

func GeneratorNum(rate float64, i int64, current *config.Auto_Inc_Struct) int64 {

	id := int64(rand.Float64() * rate) + current.ID + i * int64(rate)

	if count, err := CLotteryNum.FindId(id).Count(); err != nil {

		panic(err)

	} else {
		if count > 0 {
			id = GeneratorNum(rate, i, current)
		}
	}

	return id
}

// GeneratorNumByType根据当前的奖品数量和类型生成 中奖序列好号码
func GeneratorNumByType(ls *LotterySetting) error {

	where := bson.M{}

	where["lottery_setting_id"] = ls.ID

	_, err := CLotteryNum.RemoveAll(where)

	//current, err := config.GetAutoID(CLotteryNum.Name, MG)

	if err != nil {
		return err
	}

	giftCount := ls.GiftCount
	//rate := ls.Chance
	var i int64
	tss := make([]interface{}, giftCount)

	for i = 0; i < giftCount; i++ {

		lsn := LotterySettingNum{}
		//lsn.ID = GeneratorNum(rate, i, current)
		lsn.LotterySettingID = ls.ID
		lsn.LotteryLevel = ls.LotteryLevel
		lsn.LotteryType = ls.LotteryType

		tss[i] = &lsn

	}

	err = CLotteryNum.Insert(tss...)
	if err != nil {
		return err
	}

	return nil

}

//NewLotterySettings 批量创建LotterySetting
func NewLotterySettings(ts []*LotterySetting) ([]*LotterySetting, error) {

	var err error
	tss := make([]interface{}, len(ts))

	for i, t := range ts {

		t.ID = bson.NewObjectId()
		now := time.Now()
		t.CreatedAt = &now

		t.UpdatedAt = &now
		t.DeletedAt = &time.Time{}

		if t.LotteryType == LotteryTypeNULL {
			t.LotteryType = LotteryTypeVDOOutdoor
		}

		err = VaildDataLotterySetting(t)
		if err != nil {
			return nil, err
		}

		err := GeneratorNumByType(t) //根据type生成中奖序列

		if err != nil {
			return nil, err
		}

		if err != nil {
			return nil, err
		}
		tss[i] = t
	}

	err = CLotterySetting.Insert(tss...)

	if err != nil {
		return nil, err
	}

	return ts, nil
}

//DeleteLotterySettingsByIDs 根据数据的id数组删除数据
func DeleteLotterySettingsByIDs(ids []*bson.ObjectId) bool {

	_, err := CLotterySetting.UpdateAll(bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}, bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	})

	where := bson.M{}

	where["lottery_setting_id"] = bson.M{
		"$in": ids,
	}

	_, error := CLotteryNum.RemoveAll(where)

	if error != nil {
		err = error
	}

	return err == nil
}

//GetLotterySettingByIDs 根据数据的id数组查询数据
func GetLotterySettingByIDs(ids []*bson.ObjectId) ([]*LotterySetting, error) {
	var ts []*LotterySetting
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["_id"] = bson.M{
		"$in": ids,
	}
	err := CLotterySetting.Find(where).All(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//GetLotterySettingByIDs 根据数据的id数组查询数据
func GetLotterySettingByID(id *bson.ObjectId) (*LotterySetting, error) {
	var ls *LotterySetting
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["_id"] = id
	err := CLotterySetting.Find(where).One(&ls)
	if err != nil {

		return nil, err
	}
	return ls, nil
}

//GetLotterySettingByType 根据数据的lottery_type查询数据
func GetLotterySettingByType(ltype LotteryType) (*LotterySetting, error) {
	var ts LotterySetting
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["lottery_type"] = ltype

	err := CLotterySetting.Find(where).One(&ts)
	if err != nil {
		return nil, err
	}
	return &ts, nil
}

//GetLotterySettingsByType 根据数据的lottery_type查询数据
func GetLotterySettingsByType(ltype LotteryType, timePonit int64) ([]*LotterySetting, error) {
	var ts []*LotterySetting
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["lottery_type"] = ltype
	where["start_time"] = bson.M{
		"$lte": timePonit,
	}
	where["end_time"] = bson.M{
		"$gte": timePonit,
	}

	err := CLotterySetting.Find(where).All(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//GetLotterySettingByParams 根据数据的lottery_type查询某个时间段礼品大于0的数据
func GetLotterySettingByParams(ltype LotteryType, timePonit int64) (*LotterySetting, error) {
	var ts LotterySetting
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["lottery_type"] = ltype
	where["start_time"] = bson.M{
		"$lte": timePonit,
	}
	where["end_time"] = bson.M{
		"$gte": timePonit,
	}
	where["gift_count"] = bson.M{
		"$gt": 0,
	}

	err := CLotterySetting.Find(where).One(&ts)
	if err != nil {

		if mgo.ErrNotFound == err {

			return nil, nil
		}

		return nil, err
	}
	return &ts, nil
}

//GetLotterySettingByEndTime  根据活动类型和当前时间查找未送出的奖品
func GetLotterySettingByEndTime(ltype LotteryType, timePonit int64) (*LotterySetting, error) {
	var ts LotterySetting
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["lottery_type"] = ltype
	where["end_time"] = bson.M{
		"$lte": timePonit,
	}
	where["gift_count"] = bson.M{
		"$gt": 0,
	}

	err := CLotterySetting.Find(where).One(&ts)
	if err != nil {

		return nil, err
	}
	return &ts, nil
}

//UpdateLotterySettingsByID 批量更新LotterySetting
func UpdateLotterySettingsByID(ts []*LotterySetting) ([]*LotterySetting, error) {

	var err error
	tss := []*LotterySetting{}

	for _, m := range ts {

		now := time.Now()
		m.UpdatedAt = &now

		tss = append(tss, m)
		update := make(bson.M)
		col := make(bson.M)

		update["$set"] = col

		if !m.ID.Valid() {

			return nil, errors.New("The ID was needed,try agin!")
		}

		if m.LotteryType != LotteryTypeNULL {

			col[`lottery_type`] = m.LotteryType

		}
		if m.LotteryLevel != LotteryLevelNULL {

			col[`lottery_level`] = m.LotteryLevel

		}
		if len(strings.TrimSpace(m.LotteryName)) > 0 {

			col[`lottery_name`] = m.LotteryName

		}

		if m.GiftCount > 0 {
			col[`gift_count`] = m.GiftCount

		}
		if m.StartTime > 0 {

			col[`start_time`] = m.StartTime

		}
		if m.EndTime > 0 {

			col[`end_time`] = m.EndTime

		}
		if m.LotteryExp > 0 {

			col[`lottery_exp`] = m.LotteryExp
		}
		if m.Chance > float64(0) {

			col[`chance`] = m.Chance
		}

		col[`deleted_at`] = &time.Time{}
		col[`updated_at`] = m.UpdatedAt

		err = CLotterySetting.UpdateId(m.ID, update)

		if err != nil {
			return nil, err
		}
		if m.GiftCount > 0 || m.Chance > 0 {

			ls, err := GetLotterySettingByID(&m.ID)
			if err != nil {
				return nil, err
			}

			err = GeneratorNumByType(ls) //根据type生成中奖序列

			if err != nil {
				return nil, err
			}
		}

	}

	return tss, nil
}

//UpdateLotteryGifCountByType 更新LotterySetting 的奖品数量
func UpdateLotteryGifCountByType(gifCount int64, typen LotteryType) ([]*LotterySetting, error) {

	now := time.Now()
	update := make(bson.M)
	col := make(bson.M)

	update["$set"] = col

	if gifCount > 0 {
		col[`gift_count`] = gifCount

	}

	col[`updated_at`] = &now

	_, err := CLotterySetting.UpdateAll(bson.D{{"lottery_type", typen}}, update)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

//UpdateLotteryGifCountByID 更新LotterySetting 的奖品数量
func UpdateLotteryGifCountByID(gifCount int64, id bson.ObjectId) ([]*LotterySetting, error) {

	now := time.Now()
	update := make(bson.M)
	col := make(bson.M)

	update["$set"] = col

	col[`gift_count`] = gifCount

	col[`updated_at`] = &now

	err := CLotterySetting.Update(bson.D{{"_id", id}}, update)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

//GetLotterySettingWhereMap 根据byte数组获取map数据
func GetLotterySettingWhereMap(data []byte, where bson.M) bson.M {

	m := &LotterySetting{}

	json.Unmarshal(data, m)

	if m.LotteryType != LotteryTypeNULL {

		where[`lottery_type`] = m.LotteryType

	}
	if m.ID.Valid() {

		where[`_id`] = m.ID

	}

	return where
}

//FindLotterySettingsByParams 根据查询参数获取LotterySetting的分页数据
func FindLotterySettingsByParams(req *proto.RequestParams) ([]*LotterySetting, int64, error) {

	var ts []*LotterySetting

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

		where = GetLotterySettingWhereMap(req.Where, where)
	}

	if len(strings.TrimSpace(req.Keywords)) > 0 {

		reg := bson.M{"$regex": bson.RegEx{req.Keywords, "i"}}
		or := []bson.M{bson.M{"lottery_name": reg}}
		where["$or"] = or

	}

	query := CLotterySetting.Find(where)

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

	//totalNum mongodb中所有LotterySetting的总数，而非查询返回的条数
	totalNum, _ := CLotterySetting.Find(where).Count()

	return ts, int64(totalNum), nil
}
