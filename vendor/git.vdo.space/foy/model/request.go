package model

import (
	"encoding/json"
	"git.vdo.space/foy/proto"
)

type RequestQueryParams struct {
	StartTime  string      `json:"startTime"`
	EndTime    string      `json:"endTime"`
	OrderBy    string      `json:"orderBy"`
	IsDesc     bool        `json:"isDesc"`
	OffsetHead int64       `json:"offsetHead"`
	OffsetTail int64       `json:"offsetTail"`
	UserID     int64       `json:"userID"`
	Where      interface{} `json:"where"`
	Keywords   string      `json:"keywords"`
	UserUUID   string      `json:"user_uuid"`
}

//RequestParams2ProRequestParams 将RequestParams转换为ProRequestParams
func RequestParams2ProRequestParams(rp *RequestQueryParams) *proto.RequestParams {
	us := proto.RequestParams{}
	us.StartTime = rp.StartTime
	us.EndTime = rp.EndTime
	us.IsDesc = rp.IsDesc
	us.OffsetHead = rp.OffsetHead
	us.OffsetTail = rp.OffsetTail
	us.OrderBy = rp.OrderBy
	us.UserID = rp.UserID
	us.Keywords = rp.Keywords
	us.UserUUID = rp.UserUUID

	if rp.Where != nil {
		u, err := json.Marshal(rp.Where)
		if err == nil {
			us.Where = u
		}
	}

	return &us
}

//RequestParams2ProRequestParamsWithJson 将RequestParams转换为ProRequestParams,但是中间的二进制数据采用json转换
func RequestParams2ProRequestParamsWithJson(rp *RequestQueryParams) *proto.RequestParams {
	us := proto.RequestParams{}
	us.StartTime = rp.StartTime
	us.EndTime = rp.EndTime
	us.IsDesc = rp.IsDesc
	us.OffsetHead = rp.OffsetHead
	us.OffsetTail = rp.OffsetTail
	us.OrderBy = rp.OrderBy
	us.UserID = rp.UserID
	us.Keywords = rp.Keywords
	us.UserUUID = rp.UserUUID

	if rp.Where != nil {
		u, err := json.Marshal(rp.Where)
		if err == nil {
			us.Where = u
		}
	}

	return &us
}
