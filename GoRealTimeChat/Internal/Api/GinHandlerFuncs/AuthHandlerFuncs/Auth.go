package AuthHandlerFuncs

import (
	"GoRealTimeChat/DataModels/AuthHandlerModels"
	"GoRealTimeChat/DataModels/Internal/ApiRequests"
	"GoRealTimeChat/DataModels/Internal/ConfigModels"
	"GoRealTimeChat/DataModels/Internal/DataModels/IMUser"
	"GoRealTimeChat/DataModels/packages/Model"
	"GoRealTimeChat/HtmlTemplate/HTML/EmailCodeHtmlTemlate"
	"GoRealTimeChat/Internal/Api/Services"
	"GoRealTimeChat/Internal/Dao/AuthDao"
	"GoRealTimeChat/Internal/Enums"
	"GoRealTimeChat/Internal/Utils"
	"GoRealTimeChat/Packages/Hash"
	"GoRealTimeChat/Packages/JWT"
	"GoRealTimeChat/Packages/Response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	AUTH AuthDao.AUTHDAO
)

type AuthHandler struct {
}

// AuthHandlerInterface 定义注册登录等接口
type AuthHandlerInterface interface {
	//	Login 登录接口
	Login(ginCtx *gin.Context)
	// Registered 注册接口
	Registered(ginCtx *gin.Context)
	// SendEmailCode 发送邮件验证码接口
	SendEmailCode(ginCtx *gin.Context)
}

// Login 登录实现
func (*AuthHandler) Login(ginCtx *gin.Context) {
	// 前端提交的表单
	params := &ApiRequests.LoginForm{
		Email:    ginCtx.PostForm("email"),
		Password: ginCtx.PostForm("password"),
	}
	zap.S().Infof("login form params: %#v", params)
	// 验证表单结构
	validateError := validator.New().Struct(params)
	if validateError != nil {
		zap.S().Errorf("validate login form params error: %#v", validateError)
		Response.FailResponse(http.StatusInternalServerError, validateError.Error()).WriteTo(ginCtx)
		return
	}
	zap.S().Infof("login form params struct validate success!")
	// 查询用户是否存在
	var users IMUser.ImUsers
	result := Model.MYSQLDB.Table("im_users").Where("email=?", params.Email).First(&users)
	if result.RowsAffected == 0 {
		zap.S().Errorf("邮箱 %#v 未注册账号!", params.Email)
		Response.FailResponse(http.StatusInternalServerError, "该邮箱未注册账号").ToJson(ginCtx)
		return
	}
	//	密码错误的情况进行处理
	if !Hash.CheckPassword(params.Password, users.Password) {
		zap.S().Errorln("params中的密码错误!")
		Response.FailResponse(http.StatusInternalServerError, "密码错误").ToJson(ginCtx)
		return
	}

	// 更新JWT持续时间
	jwtTimeToLive := ConfigModels.ConfigData.JWT.TimeToLive
	expireAtTime := time.Now().Unix() + jwtTimeToLive
	token := JWT.CreateNewJWT().IssueToken(users.ID, users.Uid, users.Name, users.Email, expireAtTime)
	// 登录成功
	zap.S().Infoln("注册生成 token 成功")
	Response.SuccessResponse(&AuthHandlerModels.LoginResponse{
		ID:          users.ID,
		UID:         users.Uid,
		Name:        users.Name,
		Avatar:      users.Avatar,
		Email:       users.Email,
		Token:       token,
		ExpireTime:  expireAtTime,
		TokenToLive: jwtTimeToLive,
	}).WriteTo(ginCtx)
	return
}

// Registered 注册账号的实现
// 定义一个名为Registered的方法，属于AuthHandler结构体指针接收者
func (*AuthHandler) Registered(ginCtx *gin.Context) {
	// 创建一个ApiRequests.RegisteredForm类型的params变量，并初始化其字段值
	params := &ApiRequests.RegisteredForm{
		Email:          ginCtx.PostForm("email"),                                     // 从请求中获取email字段的值
		Name:           ginCtx.PostForm("name"),                                      // 从请求中获取name字段的值
		EmailType:      Utils.StringToInt(ginCtx.DefaultPostForm("email_type", "1")), // 从请求中获取email_type字段的值，并将其转换为整型，默认为1
		Password:       ginCtx.PostForm("password"),                                  // 从请求中获取password字段的值
		PasswordRepeat: ginCtx.PostForm("password_repeat"),                           // 从请求中获取password_repeat字段的值
		Code:           ginCtx.PostForm("code"),                                      // 从请求中获取code字段的值
	}

	// 验证 params 结构
	validateParamsStructError := validator.New().Struct(params) // 使用validator库对params进行结构验证
	if validateParamsStructError != nil {
		zap.S().Errorln(validateParamsStructError.Error())                                         // 打印验证失败的错误信息
		Response.FailResponse(Enums.ParamError, validateParamsStructError.Error()).WriteTo(ginCtx) // 返回参数错误的响应
		return
	}

	// 验证用户是否存在
	if ok, filed := Utils.IsUserExits(params.Email, params.Name); ok {
		zap.S().Errorf("%s已经存在了", filed)
		Response.FailResponse(Enums.ParamError, fmt.Sprintf("%s已经存在了", filed)).WriteTo(ginCtx) // 返回用户已经存在的响应
	}

	var emailService Services.EmailService // 创建一个EmailService类型的变量emailService

	if !emailService.CheckCode(params.Email, params.Code, params.EmailType) {
		zap.S().Errorln("邮件验证码不正确")
		Response.FailResponse(Enums.ParamError, "邮件验证码不正确").WriteTo(ginCtx) // 返回邮件验证码不正确的响应
		return
	}

	AUTH.CreateUser(params.Email, params.Password, params.Name) // 创建用户
	zap.S().Infoln("注册成功！")
	Response.SuccessResponse().ToJson(ginCtx) // 返回成功响应
	return
}

// SendEmailCode 定义一个名为SendEmailCode的方法，属于AuthHandler结构体指针接收者
func (*AuthHandler) SendEmailCode(ginCtx *gin.Context) {
	// 创建一个ApiRequests.SendEmailValidateCodeRequest类型的params变量，并初始化其字段值
	params := &ApiRequests.SendEmailValidateCodeRequest{
		Email:     ginCtx.PostForm("email"),                         // 从请求中获取email字段的值
		EmailType: Utils.StringToInt(ginCtx.PostForm("email_type")), // 从请求中获取email_type字段的值，并将其转换为整型
	}

	// 验证 params 结构
	// 使用validator库对params进行结构验证
	validateParamsStructError := validator.New().Struct(params) // 使用validator库对params进行结构验证
	if validateParamsStructError != nil {
		// 打印验证失败的错误信息
		zap.S().Errorln(validateParamsStructError.Error())
		// 返回参数错误的响应
		Response.FailResponse(Enums.ParamError, validateParamsStructError.Error()).WriteTo(ginCtx)
		return
	}
	// 检查数据库中是否存在指定的邮箱地址
	ok := ApiRequests.IsTableFliedExits("email", params.Email, "im_users")

	switch params.EmailType {
	// 如果邮箱类型为注册码
	case Services.REGISTERED_CODE:
		if ok {
			// 返回邮箱已经被注册的响应
			Response.FailResponse(Enums.ParamError, "邮箱已经被注册了").WriteTo(ginCtx)
			return
		}
		// 如果邮箱类型为重置密码码
	case Services.RESETPS_CODE:
		if !ok {
			// 返回邮箱未注册的响应
			Response.FailResponse(Enums.ParamError, "邮箱未注册").WriteTo(ginCtx)
			return
		}
	}
	// 创建一个EmailService类型的变量emailService
	var emailService Services.EmailService
	// 生成一个邮件验证码
	code := Utils.CreateEmailCode()
	// 生成邮件内容的HTML模板
	html := EmailCodeHtmlTemlate.EmailCodeHTMLTemplate(ConfigModels.ConfigData.Mail.EmailCodeHtmlTemplateFilePath, code)
	// 发送邮件
	emailServiceSendEmailError := emailService.SendEmail(code, params.EmailType, params.Email, ConfigModels.ConfigData.Mail.EmailCodeSubject, html)
	if emailServiceSendEmailError != nil {
		zap.S().Errorf("发送失败邮箱:" + params.Email + "错误日志:" + emailServiceSendEmailError.Error())
		// 返回邮件发送失败的响应
		Response.FailResponse(Enums.ApiError, "邮件发送失败,请检查是否是可用邮箱").ToJson(ginCtx)
		return
	}

	zap.S().Infoln("邮件发送成功，请注意查收！")
	// 返回成功响应
	Response.SuccessResponse().ToJson(ginCtx)
	return
}
