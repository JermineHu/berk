package consts

const (
	ResourcesPath = "ResourcesPath"
	ResourcesAddr = "ResourcesAddr"
)

const (
	Token = "Authentication-Token"
	MDBPerfix = "berk_"
	ProjectName = "berk"
)

//微服务名称
const (
	BackendGroupRouteModuleName = `/mgr`
	FrontendGroupRouteModuleName = `/api`
	ResourcesGroupRouteModuleName = `/res`
)

// 版本号
const (
	GroupRouteVersion1Key = `/v1` //版本号
)

// Controller Key
const (
	ControllerBKey = "ControllerBKey"
	ControllerFKey = "ControllerFKey"
	MeanErrorHandelKey = "ReturnErrInBodyKey"
)

// 返回码
const (
	GenericsSuccessCode = 200
)

// 用户
const (
	LoginRoute = `account/login` // POST
)

// 用户举报
const (
	GetAllTigOffParamsRoute = `tip-off` //POST
)

// 用户模块
const (
	UserListRoute = `users/list`   //POST
	UsersCreateRoute = `users`        //POST
	UsersUploadRoute = `users/upload` //POST
	UsersUpdateRoute = `users`        //PUT
	UserDetailByIDsRoute = `users/view`   //POST
	UserDeleteByIDsRoute = `users`        //Delete
)

// 用户模块
const (
	ArticleListRoute = `articles/list`   //POST
	ArticlesCreateRoute = `articles`        //POST
	ArticlesUploadRoute = `articles/upload` //POST
	ArticlesUpdateRoute = `articles`        //PUT
	ArticleDetailByIDsRoute = `articles/view`   //POST
	ArticleDeleteByIDsRoute = `articles`        //Delete
)

// 标签
const (
	TagListRoute = `tags/list` //POST
	TagsCreateRoute = `tags`      //POST
	TagsUpdateRoute = `tags`      //PUT
	TagDetailByIDsRoute = `tags/view` //POST
	TagDeleteByIDsRoute = `tags`      //Delete
)

// 表单
const (
	FormListRoute = `forms/list` //POST
	FormsCreateRoute = `forms`      //POST
	FormsUpdateRoute = `forms`      //PUT
	FormDetailByIDsRoute = `forms/view` //POST
	FormDeleteByIDsRoute = `forms`      //Delete
)

// 表单
const (
	FeedbackListRoute = `feedbacks/list` //POST
	FeedbacksCreateRoute = `feedbacks`      //POST
	FeedbacksUpdateRoute = `feedbacks`      //PUT
	FeedbackDetailByIDsRoute = `feedbacks/view` //POST
	FeedbackDeleteByIDsRoute = `feedbacks`      //Delete
)

// 微信接口
const (
	GetWXConfigRoute = `wx_config`       //GET
	WXOauth2Route = `wx_Oauth2/:Code` //GET

)

// 活动模块
const (
	ActivityListRoute = `activities/list`   //POST
	ActivitysCreateRoute = `activities`        //POST
	ActivitysUploadRoute = `activities/upload` //POST
	ActivitysUpdateRoute = `activities`        //PUT
	ActivityDetailByIDsRoute = `activities/view`   //POST
	ActivityDeleteByIDsRoute = `activities`        //Delete
)


// 抽奖设置模块
const (
	LotterySettingListRoute = `lottery/list`      //POST
	LotteryGeneratorListRoute = `lottery/generator` //POST
	LotterySettingsCreateRoute = `lottery`           //POST
	LotterySettingsUpdateRoute = `lottery`           //PUT
	LotterySettingDetailByIDsRoute = `lottery/view`      //POST
	LotterySettingDeleteByIDsRoute = `lottery`           //Delete
)

// 抽奖结果模块
const (
	LotteryDrawIndoorRoute = `lottery/draw_indoor`  //GET
	LotteryDraw20170610Route = `lottery/draw_20170610`  //GET
	LotteryDrawOutdoorRoute = `lottery/draw_outdoor` //GET
	LotteryDrawListRoute = `winning/list`         //POST
	LotteryDrawsCreateRoute = `winning`              //POST
	LotteryDrawsUpdateAwardRoute = `winning/award`        //PUT
	LotteryDrawsUpdateRoute = `winning`              //PUT
	LotteryDrawDetailByIDsRoute = `winning/view`         //POST
	LotteryDrawDetailByUserUUIDIndoorRoute = `winning/indoor`       //GET
	LotteryDrawDetailByUserUUIDOutdoorRoute = `winning/outdoor`      //GET
	LotteryDrawDetailByUserUUIDHamleys20170610Route = `winning/hamleys20170610`      //GET
	LotteryDrawIndoorStatusRoute = `indoor/status`        //GET
	LotteryDrawOutdoorStatusRoute = `outdoor/status`       //GET
	LotteryDrawDeleteByIDsRoute = `winning`              //Delete
)

var ImageFileExts = []string{".png", ".jpeg", ".jpg", ".bmp", ".gif"}

var FilterFileExts = []string{".3dl", ".cube", ".png"}
