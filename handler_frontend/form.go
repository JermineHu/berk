package handler_frontend

import (
	"git.vdo.space/berk/model"
	dm "git.vdo.space/foy/model"
)

//FormCreate 创建标签
func (m *FController) FormsCreate(ts []*model.Form) (*dm.ResultModel, error) {

	rm := dm.ResultModel{}
	for i, _ := range ts {
		ts[i].UserID = m.User.ID
	}
	t, err := model.NewForms(ts)

	if err != nil {
		rm.Errors = append(rm.Errors, dm.NewError(err))
	} else {
		rm.Data = t
	}

	return &rm, nil
}
