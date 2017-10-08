package handler_backend

import (
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"git.vdo.space/berk/model"
	"gopkg.in/mgo.v2/bson"
)

//LotterySettingCreate 创建LotterySetting
func (m *BController) LotterySettingsCreate(ts []*model.LotterySetting) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ms, err := model.NewLotterySettings(ts)

	if err != nil {
		return nil, err
	} else {
		d, err := model.LotterySettingsToDatas(ms)
		if err != nil {
			return nil, err
		}
		rm.Data = d
	}

	return &rm, nil
}

//LotterySettingUpdate 修改LotterySetting
func (m *BController) LotterySettingsUpdate(imgf []*model.LotterySetting) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ts, err := model.UpdateLotterySettingsByID(imgf)

	if err != nil {
		return nil, err
	} else {

		d, err := model.LotterySettingsToDatas(ts)
		if err != nil {
			return nil, err
		}
		rm.Data = d
	}

	return &rm, nil
}

//LotterySettingList LotterySetting列表
func (m *BController) LotterySettingList(req *proto.RequestParams) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	list, count, err := model.FindLotterySettingsByParams(req)

	if err != nil {

		panic(dm.NewError(err))
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {

		d := []dm.Dungeons{}

		for _, item := range list {

			dg, err := model.LotterySettingToData(item)

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

//LotterySettingDetailByIDs LotterySetting根据id数组获取LotterySetting详情
func (m *BController) LotterySettingDetailByIDs(ids []*bson.ObjectId) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}
	ts, err := model.GetLotterySettingByIDs(ids)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		d, err := model.LotterySettingsToDatas(ts)
		if err != nil {
			return nil, err
		}
		rm.Data = d
	}

	return &rm, nil
}

//LotterySettingDetailByIDs LotterySetting根据id数组获取LotterySetting详情
func (m *BController) LotterySettingDeleteByIDs(ids []*bson.ObjectId) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	d := make(dm.Dungeons)

	d["success"] = model.DeleteLotterySettingsByIDs(ids)

	rm.Data = d

	return &rm, nil
}
