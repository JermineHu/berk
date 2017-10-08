package handler_backend

import (
	dm "git.vdo.space/foy/model"
	"git.vdo.space/foy/proto"
	"git.vdo.space/berk/model"
	"qiniupkg.com/x/errors.v7"
)

//LotteryDrawCreate 创建LotteryDraw
func (m *BController) LotteryDrawsCreate(ts []*model.LotteryDraw) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	ms, err := model.NewLotterys(ts)

	if err != nil {
		return nil, err
	} else {
		d, err := model.LotteryDrawsToDatas(ms)
		if err != nil {
			return nil, err
		}
		rm.Data = d
	}

	return &rm, nil
}

//LotteryDrawUpdate 修改LotteryDraw
func (m *BController) LotteryDrawsUpdate(imgf []*model.LotteryDraw) (*dm.ResultModel, error) {

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

//LotteryDrawUpdate 修改LotteryDraw
func (m *BController) LotteryDrawsUpdateAward(ids []int64) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	count, err := model.GetLotteryAwardStatusByIDs(ids)

	if err != nil {
		panic(dm.NewError(err))
	}

	if count > 0 {
		panic(dm.NewError(errors.New("The lottery has been Award!")))
	}
	d := make(dm.Dungeons)

	d["success"] = model.LotteryDrawsUpdateAwardByIDs(ids)

	rm.Data = d

	return &rm, nil
}

//LotteryDrawList LotteryDraw列表
func (m *BController) LotteryDrawList(req *proto.RequestParams) (*dm.ResultModel, error) {

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
func (m *BController) LotteryDrawDetailByIDs(ids []int64) (*dm.ResultModel, error) {

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

//LotteryDrawDetailByIDs LotteryDraw根据id数组获取LotteryDraw详情
func (m *BController) LotteryDrawDeleteByIDs(ids []int64) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}

	d := make(dm.Dungeons)

	d["success"] = model.DeleteLotterysByIDs(ids)

	rm.Data = d

	return &rm, nil
}
