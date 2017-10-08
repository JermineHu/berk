package model

import (
	"encoding/json"
	"errors"
	. "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"strings"
	"time"
)

type LotteryStatus int32

const (
	LotteryStatusNULL LotteryStatus = 0
	LotteryStatusAward LotteryStatus = 1
	LotteryStatusNotAward LotteryStatus = 2
)

//GetStrByLotteryStatus 获取状态的string
func GetStrByLotteryStatus(status LotteryStatus) string {

	switch status {

	case LotteryStatusAward:

		return "LotteryStatusAward"

		break

	case LotteryStatusNotAward:

		return "LotteryStatusNotAward"

		break
	}

	return "LotteryStatusNULL"

}

//GetLotteryStatusByStr 根据string获取状态
func GetLotteryStatusByStr(status string) LotteryStatus {

	switch status {

	case "LotteryStatusAward":

		return LotteryStatusAward

		break

	case "LotteryStatusNotAward":

		return LotteryStatusNotAward

		break
	}

	return LotteryStatusNULL

}

type WinningStatus int32

const (
	WinningStatusNULL WinningStatus = 0
	WinningStatusNotWinning WinningStatus = 1
	WinningStatusWinning WinningStatus = 2
)

//GetStrByWinningStatus 获取状态的string
func GetStrByWinningStatus(status WinningStatus) string {

	switch status {

	case WinningStatusNotWinning:

		return "WinningStatusNotWinning"

		break

	case WinningStatusWinning:

		return "WinningStatusWinning"

		break
	}

	return "WinningStatusNULL"

}

//GetWinningStatusByStr 根据string获取状态
func GetWinningStatusByStr(status string) WinningStatus {

	switch status {

	case "WinningStatusNotWinning":

		return WinningStatusNotWinning

		break

	case "WinningStatusWinning":

		return WinningStatusWinning

		break
	}

	return WinningStatusNULL

}

type LotteryDraw struct {
	ID            int64         `bson:"id" json:"id"`                                             //数据自增编号
	Number        string        `json:"number,omitempty" bson:"number,omitempty"`                 //抽奖编号
	WinningStatus WinningStatus `bson:"winning_status,omitempty" json:"winning_status,omitempty"` // 中奖没有
	LotteryStatus LotteryStatus `bson:"lottery_status,omitempty" json:"lottery_status,omitempty"` //领过奖没有
	LotteryType   LotteryType   `bson:"lottery_type,omitempty" json:"lottery_type,omitempty"`     // 抽奖类型
	LotteryLevel  LotteryLevel  `bson:"lottery_level,omitempty" json:"lottery_level"`             // 获奖级别
	UserUUID      *bson.ObjectId        `json:"user_uuid,omitempty" bson:"user_uuid,omitempty"`
	LotterySettingID bson.ObjectId `bson:"lottery_setting_id" json:"lottery_setting_id"`
	AwardAt       *time.Time    `bson:"award_at"   json:"award_at"`                               //领奖时间
	CreatedAt     *time.Time    `bson:"created_at" json:"created_at"`
	UpdatedAt     *time.Time    `bson:"updated_at" json:"updated_at"`
	DeletedAt     *time.Time    `bson:"deleted_at" json:"deleted_at"`
}

func (v LotteryDraw) GetCollectionName() string {

	return reflect.TypeOf(v).Name()
}

func LotteryDrawToData(md *LotteryDraw) (Dungeons, error) {
	d := make(Dungeons)

	if md != nil {
		d[`id`] = md.ID
		d[`number`] = md.Number
		d[`lottery_status`] = GetStrByLotteryStatus(md.LotteryStatus)
		d[`winning_status`] = GetStrByWinningStatus(md.WinningStatus)
		d[`lottery_type`] = GetStrByLotteryType(md.LotteryType)
		d[`lottery_level`] = GetStrByLotteryLevel(md.LotteryLevel)
		d[`user_uuid`] = md.UserUUID
		d[`award_at`] = md.AwardAt
		d[`created_at`] = md.CreatedAt
		d[`updated_at`] = md.UpdatedAt
	}

	return d, nil
}

func LotteryDrawsToDatas(mds []*LotteryDraw) ([]Dungeons, error) {
	dl := []Dungeons{}
	for _, x := range mds {
		d, err := LotteryDrawToData(x)
		if err != nil {
			return nil, err
		}
		dl = append(dl, d)
	}

	return dl, nil
}

// VaildDataImageFilter 验证数据
func VaildDataLottery(t *LotteryDraw) error {

	if len(strings.TrimSpace(t.Number)) == 0 {

		return errors.New("The Number does not a null string,try agin!")
	}
	if !t.UserUUID.Valid() {

		return errors.New("The UserUUID does not a null string,try agin!")
	}
	if t.ID == 0 {

		return errors.New("The id was required,try agin!")
	}

	if t.LotteryType == LotteryTypeNULL {

		return errors.New("The LotteryType was needed,try agin!")
	}
	if t.ID == 0 {

		return errors.New("The ID was needed,try agin!")
	}

	return nil

}

//NewLotterys 批量创建Lottery
func NewLotterys(ts []*LotteryDraw) ([]*LotteryDraw, error) {

	var err error
	tss := make([]interface{}, len(ts))

	for i, t := range ts {

		now := time.Now()
		t.CreatedAt = &now

		t.UpdatedAt = &now
		t.DeletedAt = &time.Time{}

		if t.LotteryStatus == LotteryStatusNULL {
			t.LotteryStatus = LotteryStatusNotAward
		}
		if t.WinningStatus == WinningStatusNULL {
			t.WinningStatus = WinningStatusNotWinning
		}

		err = VaildDataLottery(t)
		if err != nil {
			return nil, err
		}
		tss[i] = t

	}

	err = CLottery.Insert(tss...)

	if err != nil {
		return nil, err
	}

	return ts, nil
}

//DeleteLotterysByIDs 根据数据的id数组删除数据
func DeleteLotterysByIDs(ids []int64) bool {

	_, err := CLottery.UpdateAll(bson.M{
		"id": bson.M{
			"$in": ids,
		},
	}, bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	})

	return err == nil
}

//LotteryDrawsUpdateAwardByIDs 根据数据的id数组批量设置领奖状态数据
func LotteryDrawsUpdateAwardByIDs(ids []int64) bool {

	_, err := CLottery.UpdateAll(bson.M{
		"id": bson.M{
			"$in": ids,
		},
	}, bson.M{
		"$set": bson.M{
			"lottery_status": LotteryStatusAward,
			"award_at": time.Now(),
		},
	})

	return err == nil
}

//GetLotteryByIDs 根据数据的id数组查询数据
func GetLotteryByIDs(ids []int64) ([]*LotteryDraw, error) {
	var ts []*LotteryDraw
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["id"] = bson.M{
		"$in": ids,
	}
	err := CLottery.Find(where).All(&ts)
	if err != nil {

		return nil, err
	}
	return ts, nil
}

//GetLotteryAwardStatusByIDs 根据数据的id数组查询数据
func GetLotteryAwardStatusByIDs(ids []int64) (int, error) {
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["id"] = bson.M{
		"$in": ids,
	}

	where["lottery_status"] = LotteryStatusAward

	count, err := CLottery.Find(where).Count()
	if err != nil {

		return 0, err
	}
	return count, nil
}

//GetLotteryByID 根据数据的id查询数据
func GetLotteryByUserUUID(uuid *bson.ObjectId, lt LotteryType) (*LotteryDraw, error) {
	var ts LotteryDraw
	where := bson.M{}
	where["deleted_at"] = bson.M{
		"$lte": time.Time{},
	}
	where["user_uuid"] = uuid
	if lt != LotteryTypeNULL {
		where["lottery_type"] = lt
	}
	err := CLottery.Find(where).One(&ts)
	if err != nil {

		if err == mgo.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &ts, nil
}

//UpdateLotterysByID 批量更新Lottery
func UpdateLotterysByID(ts []*LotteryDraw) ([]*LotteryDraw, error) {

	var err error
	tss := []*LotteryDraw{}

	for _, m := range ts {

		now := time.Now()
		m.UpdatedAt = &now

		tss = append(tss, m)
		update := make(bson.M)
		col := make(bson.M)

		update["$set"] = col

		if m.LotteryStatus != LotteryStatusNULL {

			col[`lottery_status`] = m.LotteryStatus

		}
		if m.WinningStatus != WinningStatusNULL {

			col[`winning_status`] = m.WinningStatus

		}
		if m.LotteryType != LotteryTypeNULL {

			col[`lottery_type`] = m.LotteryType

		}
		if m.UserUUID.Valid() {

			col[`user_uuid`] = m.UserUUID
		}
		col[`deleted_at`] = &time.Time{}
		col[`updated_at`] = m.UpdatedAt

		err = CLottery.Update(bson.D{{"id", m.ID}}, update)

		if err != nil {
			return nil, err
		}

	}

	return tss, nil
}

//GetLotteryWhereMap 根据byte数组获取map数据
func GetLotteryWhereMap(data []byte, where bson.M) bson.M {

	m := &LotteryDraw{}

	json.Unmarshal(data, m)

	if m.LotteryStatus != LotteryStatusNULL {

		where[`lottery_status`] = m.LotteryStatus

	}
	if m.WinningStatus != WinningStatusNULL {

		where[`winning_status`] = m.WinningStatus

	}
	if m.LotteryType != LotteryTypeNULL {

		where[`lottery_type`] = m.LotteryType

	}
	if m.UserUUID.Valid() {

		where[`user_uuid`] = m.UserUUID
	}

	if m.ID != 0 {

		where[`id`] = m.ID

	}

	if len(strings.TrimSpace(m.Number)) > 0 {

		where["number"] = m.Number
	}

	return where
}

//FindLotterysByParams 根据查询参数获取Lottery的分页数据
func FindLotterysByParams(req *proto.RequestParams) ([]*LotteryDraw, int64, error) {

	var ts []*LotteryDraw

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

		where = GetLotteryWhereMap(req.Where, where)
	}

	if len(strings.TrimSpace(req.Keywords)) > 0 {

		reg := bson.M{"$regex": bson.RegEx{req.Keywords, "i"}}
		or := []bson.M{bson.M{"number": reg}}
		where["$or"] = or

	}

	query := CLottery.Find(where)

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

	//totalNum mongodb中所有Lottery的总数，而非查询返回的条数
	totalNum, _ := CLottery.Find(where).Count()

	return ts, int64(totalNum), nil
}
