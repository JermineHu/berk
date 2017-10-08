package common

import (
	"encoding/json"
	"git.vdo.space/foy/model"
)

type PageModelBase struct {
	Count    int64       `json:"count,omitempty"`
	PageData interface{} `json:"list,omitempty"`
}

// params model that server return data and message info
type ResultModel struct {
	Errors []*model.ErrorMsgDefault `json:"errors,omitempty"`
	Data   interface{}              `json:"data,omitempty"`
}

func (p ResultModel) Error() string {

	str, _ := json.Marshal(p.Errors)

	return string(str)
}
