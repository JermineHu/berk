package handler_frontend

import (
	"errors"
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"git.vdo.space/berk/model"
	"gopkg.in/mgo.v2/bson"
	"time"
	"github.com/encryptio/alias"
	"math/rand"
	"fmt"
	"git.vdo.space/foy/config"
	"gopkg.in/mgo.v2"
)

/**
TODO：
 1、检查是否登陆过；
 2、查找当前正在进行的活动，没开始提示未开始，活动结束就提示活动结束；
 3、检查是否参与过；
 4、抽奖算法；
 5、记录中奖结果，中奖后需要查询当前时间段内的奖品送出数量，如果当前中奖人数超过设置送出数量则设置为未中奖，如果当前时间段内无人中奖，则会在最后结束的一个小时内送完为止；
 6、中奖后奖品的总数减1；

*/
//LotteryDrawOutdoorCreate 创建LotteryDraw的outdoor
func (m *FController) LotteryDrawActivityCreate(ld model.LotteryDraw) (*dm.ResultModel, error) {

	// TODO 1 、根据活动id先判定活动的时间段，活动的状态
	// TODO 2 、检查否登陆
	// TODO 2 、检查是否参与过
	// TODO 2 、根据中奖序列获取用户中奖情况

	// GetActivityByID()


	 ls, err := model.GetLotterySettingByType(ld.LotteryType) //根据类型获取设置

	if err != nil {

		return nil, err
	}

	now := time.Now()

	et := time.Unix(ls.EndTime, 0)
	st := time.Unix(ls.StartTime, 0)

	if now.After(et) {
		//活动结束

		return nil, dm.NewError(errors.New("The lottery activity has been end !"))
	}

	if now.Before(st) {
		//活动未开始

		return nil, dm.NewError(errors.New("The lottery activity does not start !"))
	}

	ln, err := model.GetThirdPartyLoginByUserIDAndType(m.User.ID, model.LoginType_WECHAT)

	if err != nil {

		return nil, err
	}

	if ln == nil {
		return nil, errors.New(" Need to join this activity !")
	}

	ld.Number = ln.LotteryDrawNum
	_, err = m.GetParticipatedActivityByLotteryType(ld.LotteryType)

	if err != nil {

		return nil, err
	}


	//奖品数量大于0

	current, _ := config.GetAutoIncID(model.CLotteryNum.Name, model.MG)

	where := bson.M{}
	where["_id"] = current.ID
	where["lottery_type"] = ld.LotteryType
	ld.ID = current.ID
	ld.LotterySettingID = ls.ID

	count, err := model.CLotteryNum.Find(where).Count()

	if err != nil && err != mgo.ErrNotFound {
		return nil, err
	}

	if count > 0 {

		ld.WinningStatus = model.WinningStatusWinning
		ld.LotteryLevel = ls.LotteryLevel

	}



	rm := dm.ResultModel{}
	ts := []*model.LotteryDraw{&ld}
	ms, err := model.NewLotterys(ts)

	if err != nil {
		return nil, err
	} else {

		if ld.WinningStatus == model.WinningStatusWinning {
			//中奖后礼品要减1
			ls.GiftCount--
			_, err := model.UpdateLotteryGifCountByID(ls.GiftCount, ld.LotterySettingID)

			if err != nil {
				return nil, err
			}

		}

		d, err := model.LotteryDrawToData(ms[0])
		if err != nil {
			return nil, err
		}
		rm.Data = d
	}

	return &rm, nil
}

// DistributionResult 抽奖算法； 采用Alias Method 巧妙实现
func DistributionResult(prob []float64, seed int64) uint32 {

	alias, err := alias.New(prob)

	if err != nil {
		panic(err)
	}
	//seed:=time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	return alias.Gen(rng)

}

// LotteryDraw 抽奖算法； 采用Alias Method 巧妙实现
func LotteryDrawRegularByAlias() {

	prize := []model.LotteryLevel{model.LotteryLevelFirst, model.LotteryLevelSecond, model.LotteryLevelNULL}

	prob := []float64{30, 35, 35}

	index := DistributionResult(prob, time.Now().UTC().UnixNano())

	level := model.GetStrByLotteryLevel(prize[index])

	fmt.Println("--->", index, level)

	/**
	TODO
	 1、中奖后需要查询当前时间段内的奖品送出数量，如果当前中奖人数超过设置送出数量则设置为未中奖
	 2、如果当前时间段内无人中奖，则会在最后结束的一个小时内送完为止
	 */


}

// LotteryDrawRegularByTimeRange 抽奖算法； 按照时间段设定抽奖,每个时间段的第一个参与的就中奖
func LotteryDrawRegularByTimeRange(ltype model.LotteryType, ld model.LotteryDraw) (model.LotteryDraw, *model.LotterySetting) {

	ls, err := model.GetLotterySettingByParams(ltype, time.Now().Unix()) // 根据活动类型和当前时间查找当前进行中的活动

	if err != nil {

		panic(err)
	}

	if ls != nil {
		ld.WinningStatus = model.WinningStatusWinning
		ld.LotteryLevel = ls.LotteryLevel
		return ld, ls

	}

	// 如果在最后的1个小时中，发现还有没有送完的礼品，就自动随机抽取，直到送完：

	prize := []model.LotteryLevel{model.LotteryLevelFirst, model.LotteryLevelSecond, model.LotteryLevelNULL}

	prob := []float64{10, 10, 80}

	index := DistributionResult(prob, time.Now().UTC().UnixNano())

	if prize[index] == model.LotteryLevelFirst || prize[index] == model.LotteryLevelSecond {

		ls, err := model.GetLotterySettingByEndTime(ltype, time.Now().Unix()) // 根据活动类型和当前时间查找未送出的奖品

		if err != nil {

			panic(err)
		}

		if ls != nil {
			ld.WinningStatus = model.WinningStatusWinning
			ld.LotteryLevel = ls.LotteryLevel
			return ld, ls

		}

	}

	return ld, ls

}

/**
TODO：
 1、检查是否登陆过；
 2、查找当前正在进行的活动，没开始提示未开始，活动结束就提示活动结束；
 3、检查是否参与过；
 4、抽奖算法；
 5、记录中奖结果，中奖后需要查询当前时间段内的奖品送出数量，如果当前中奖人数超过设置送出数量则设置为未中奖，如果当前时间段内无人中奖，则会在最后结束的一个小时内送完为止；
 6、中奖后奖品的总数减1；

*/
//LotteryDrawCreate 创建LotteryDraw的outdoor
func (m *FController) LotteryDrawCreate(ltype model.LotteryType, activitySt, activityEt int64) (*dm.ResultModel, error) {

	ld := model.LotteryDraw{}
	ld.LotteryType = ltype
	ld.WinningStatus = model.WinningStatusNotWinning
	ld.UserUUID = m.User.ID

	ln, err := model.GetThirdPartyLoginByUserIDAndType(m.User.ID, model.LoginType_WECHAT) // 检查是否登陆过

	if err != nil {

		return nil, err
	}

	if ln == nil {
		return nil, errors.New(" Need to login !")
	}

	now := time.Now()  // 根据当前时间判断活动是否开始

	et := time.Unix(activityEt, 0)
	st := time.Unix(activitySt, 0)

	if now.After(et) {
		//活动结束

		return nil, dm.NewError(errors.New("The lottery activity has been end !"))
	}

	if now.Before(st) {
		//活动未开始

		return nil, dm.NewError(errors.New("The lottery activity does not start !"))
	}

	ld.Number = ln.LotteryDrawNum
	_, err = m.GetParticipatedActivityByLotteryType(ltype) //检查是否已经参与过了

	if err != nil {

		return nil, err
	}

	ld, lsm := LotteryDrawRegularByTimeRange(ltype, ld) // 抽奖,并设置抽奖状态


	rm := dm.ResultModel{}
	ts := []*model.LotteryDraw{&ld}

	ms, err := model.NewLotterys(ts) //记录中奖结果

	if err != nil {
		return nil, err
	} else {

		if ld.WinningStatus == model.WinningStatusWinning {
			//中奖后礼品要减1
			lsm.GiftCount--
			_, err := model.UpdateLotteryGifCountByID(lsm.GiftCount, lsm.ID)

			if err != nil {
				return nil, err
			}

		}

		d, err := model.LotteryDrawToData(ms[0])
		if err != nil {
			return nil, err
		}
		rm.Data = d
	}

	return &rm, nil
}

//LotteryDrawUpdate 修改LotteryDraw
func (m *FController) LotteryDrawsUpdate(imgf []*model.LotteryDraw) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ts, err := model.UpdateLotterysByID(imgf)

	if err != nil {
		return nil, err
	} else {

		d, err := model.LotteryDrawsToDatas(ts)
		if err != nil {
			return nil, err
		}
		rm.Data = d
	}

	return &rm, nil
}

//LotteryDrawList LotteryDraw列表
func (m *FController) LotteryDrawList(req *proto.RequestParams) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	list, count, err := model.FindLotterysByParams(req)

	if err != nil {

		panic(dm.NewError(err))
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {

		d := []dm.Dungeons{}

		for _, item := range list {

			dg, err := model.LotteryDrawToData(item)

			if err != nil {

				rm.Errors = append(rm.Errors, dm.NewError(err))
			}

			d = append(d, dg)

		}

		pm := dm.PageModelBase{}
		pm.PageData = d
		pm.Count = count
		rm.Data = pm
	}

	return &rm, nil
}

//LotteryDrawDetailByIDs LotteryDraw根据id数组获取LotteryDraw详情
func (m *FController) LotteryDrawDetailByIDs(ids []int64) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}
	ts, err := model.GetLotteryByIDs(ids)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		d, err := model.LotteryDrawsToDatas(ts)
		if err != nil {
			return nil, err
		}
		rm.Data = d
	}

	return &rm, nil
}

//LotteryDrawDetailByUserUUIDAndType LotteryDraw根据id获取LotteryDraw详情
func (m *FController) LotteryDrawDetailByUserUUIDAndType(Ltype model.LotteryType) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}
	ts, err := model.GetLotteryByUserUUID(m.User.ID, Ltype)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		d, err := model.LotteryDrawToData(ts)
		if err != nil {
			return nil, err
		}
		rm.Data = d
	}

	return &rm, nil
}

//GetLotteryDrawStatusIndoor 获取活动的开始结束状态
func (m *FController) GetLotteryDrawStatusIndoor() (*dm.ResultModel, error) {

	rm := dm.ResultModel{}
	d := make(dm.Dungeons)

	ls, err := model.GetLotterySettingByType(model.LotteryTypeVDOIndoor) //根据类型获取设置

	if err != nil {

		return nil, err
	}

	now := time.Now()

	et := time.Unix(ls.EndTime, 0)
	st := time.Unix(ls.StartTime, 0)

	if now.After(et) {
		//活动结束

		d["status"] = "Ended"
	}

	if now.Before(st) {
		//活动未开始

		d["status"] = "NotStart"
	}

	rm.Data = d
	return &rm, nil
}

//GetLotteryDrawStatusOutdoor  获取活动的开始结束状态
func (m *FController) GetLotteryDrawStatusOutdoor() (*dm.ResultModel, error) {

	rm := dm.ResultModel{}
	d := make(dm.Dungeons)
	ls, err := model.GetLotterySettingByType(model.LotteryTypeVDOOutdoor) //根据类型获取设置

	if err != nil {

		return nil, err
	}

	now := time.Now()
	et := time.Unix(ls.EndTime, 0)
	st := time.Unix(ls.StartTime, 0)

	if now.After(et) {
		//活动结束

		d["status"] = "Ended"
	}

	if now.Before(st) {
		//活动未开始

		d["status"] = "NotStart"
	}

	rm.Data = d
	return &rm, nil
}

//LotteryDrawDetailByIDs LotteryDraw根据id数组获取LotteryDraw详情
func (m *FController) LotteryDrawDeleteByIDs(ids []int64) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}
	d := make(dm.Dungeons)
	d["success"] = model.DeleteLotterysByIDs(ids)
	rm.Data = d
	return &rm, nil
}

//GetParticipatedActivityByLotteryType 是参与过活动
func (m *FController) GetParticipatedActivityByLotteryType(ltype model.LotteryType) (*model.LotteryDraw, error) {

	mld, err := model.GetLotteryByUserUUID(m.User.ID, ltype)

	if err != nil {
		return nil, err
	}

	if mld != nil {
		//不为nil说明已经参加过了
		return nil, errors.New("This activities have participated !")
	}

	return mld, nil
}
