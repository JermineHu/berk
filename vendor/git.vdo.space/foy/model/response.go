package model

import (
	"encoding/json"
	"fmt"
	"git.vdo.space/foy/utils"
)

//TODO 错误的代码可能不需要几个 之后可以合并
type ErrorMsgDefault struct {
	HttpCode  int             `json:"-"`
	ErrorType utils.ErrorType `json:"-"`
	ErrorCode int             `json:"code,omitempty"`
	Message   string          `json:"message,omitempty"`
}

type PageModelBase struct {
	Count    int64       `json:"count,omitempty"`
	PageData interface{} `json:"list,omitempty"`
}

// params model that server return data and message info
type ResultModel struct {
	Errors []*ErrorMsgDefault `json:"errors,omitempty"`
	Data   interface{}        `json:"data,omitempty"`
}

func (p ErrorMsgDefault) Error() string {
	return fmt.Sprintf("Code: " + string(int(p.ErrorCode)) + ", Message: " + p.Message)
}

func (p ResultModel) Error() string {

	str, _ := json.Marshal(p.Errors)

	return string(str)
}

func NewReultModel(data interface{}, e ...*ErrorMsgDefault) ResultModel {

	return ResultModel{
		Errors: e,
		Data:   data,
	}
}

func NewResponseError(message string, errorCode, httpCode int, errorType utils.ErrorType) *ErrorMsgDefault {
	return &ErrorMsgDefault{
		Message:   message,
		ErrorCode: errorCode,
		HttpCode:  httpCode,
		ErrorType: errorType,
	}
}

func NewError(err error) *ErrorMsgDefault {
	return &ErrorMsgDefault{
		Message:   err.Error(),
		ErrorCode: 1,
		HttpCode:  200,
		ErrorType: 1,
	}
}

func SuccessResp(content string, data interface{}) Dungeons {
	var d Dungeons
	d[`code`] = 0
	d[`message`] = content
	d[`data`] = data
	return d
}
