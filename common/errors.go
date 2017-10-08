package common

import (
	"git.vdo.space/foy/model"
	"git.vdo.space/foy/utils"
)

var (
	ServerErr                  = model.NewResponseError("ServerErr", -1, 200, utils.ErrorTypeInternal)
	RequestParamsErr           = model.NewResponseError("Request parameter invalid", 999, 200, utils.ErrorTypePublic)
	StringConversionErr        = model.NewResponseError("string conversion invalid", 999, 400, utils.ErrorTypePublic)
	AuthErr                    = model.NewResponseError("Un Auth Access", 1, 401, utils.ErrorTypePublic)
	InsufficientPermissionsErr = model.NewResponseError("Insufficient permissions", 2, 403, utils.ErrorTypePublic)
	DeviceErr                  = model.NewResponseError("Device Token Error", 3, 401, utils.ErrorTypePublic)
	ServiceErr                 = model.NewResponseError("Servicer error", 500, 500, utils.ErrorTypePublic)
)

//account
var (
	UserNotFoundErr       = model.NewResponseError("user not found", 1000, 200, utils.ErrorTypePublic)
	InvalidEmailAdressErr = model.NewResponseError("Invalid Email", 1001, 200, utils.ErrorTypePublic)
	InvalidPhoneNumberErr = model.NewResponseError("Invalid Phone Number", 1002, 200, utils.ErrorTypePublic)
	InvalidUserNameErr    = model.NewResponseError("Invalid User name", 1003, 200, utils.ErrorTypePublic)
	InvalidPasswordErr    = model.NewResponseError("Wrong Password", 1004, 200, utils.ErrorTypePublic)
	NameFormatErr         = model.NewResponseError("name format error", 1005, 200, utils.ErrorTypePublic)
	InvalidSecretCodeErr  = model.NewResponseError("secret code error", 1006, 200, utils.ErrorTypePublic)
	TimestampFormatErr    = model.NewResponseError("time stamp format error", 1007, 200, utils.ErrorTypePublic)
	AuthenticationErr     = model.NewResponseError("Authentication failure!", 1008, 200, utils.ErrorTypePublic)
)
