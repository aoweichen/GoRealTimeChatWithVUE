package Middleware

import (
	"GoRealTimeChat/Packages/JWT"
	"GoRealTimeChat/Packages/Response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// Auth 定义 Auth 函数，返回一个 gin.HandlerFunc
func Auth() gin.HandlerFunc {
	// 返回一个匿名函数作为中间件函数
	return func(ginCtx *gin.Context) {
		// 从请求的 URL 参数和请求头中获取 token
		token := ginCtx.DefaultQuery("token", ginCtx.GetHeader("authorization"))
		// 调用 ValidatedToken 函数验证 token 的有效性，返回验证错误和处理后的 token
		validatedTokenError, token := ValidatedToken(token)
		// 如果验证错误不为空，则表示 token 无效
		if validatedTokenError != nil {
			// 使用 zap 记录错误日志
			zap.S().Errorln(validatedTokenError)
			// 返回未授权的错误响应，并设置 HTTP 状态码为 401
			Response.ErrorResponse(http.StatusUnauthorized, validatedTokenError.Error()).SetHttpCode(http.StatusUnauthorized).WriteTo(ginCtx)
			// 终止后续中间件和请求处理函数的执行
			ginCtx.Abort()
			return
		}
		// 调用 CreateNewJWT 函数创建一个新的 JWT 实例，然后调用 ParseToken 方法解析 token，并返回解析后的 claims 和解析错误
		claims, jWTCreateNewJWTParseTokenError := JWT.CreateNewJWT().ParseToken(token)
		// 如果解析错误不为空，则表示 token 解析失败
		if jWTCreateNewJWTParseTokenError != nil {
			// 使用 zap 记录错误日志
			zap.S().Errorln(jWTCreateNewJWTParseTokenError)
			// 返回未授权的错误响应，并设置 HTTP 状态码为 401
			Response.ErrorResponse(http.StatusUnauthorized, jWTCreateNewJWTParseTokenError.Error()).SetHttpCode(http.StatusUnauthorized).WriteTo(ginCtx)
			// 终止后续中间件和请求处理函数的执行
			ginCtx.Abort()
			return
		}
		// 将 claims 中的 ID、UID、Name 设置到 gin 的上下文中，供后续的处理函数使用
		ginCtx.Set("id", claims.ID)
		fmt.Println("claims ", claims)
		ginCtx.Set("uid", claims.UID)
		ginCtx.Set("name", claims.Name)
	}
}

func ValidatedToken(token string) (error, string) {
	var err error // 声明一个error类型的变量err

	if len(token) == 0 {
		err = errors.New("Token不能为空")
		return err, err.Error() // 如果token为空，则返回err和"Token 不能为空"的错误信息
	}

	t := strings.Split(token, "Bearer ") // 使用空格分割token字符串，并将结果赋值给t变量
	if len(t) > 1 {
		return nil, t[1] // 如果t的长度大于1，则返回nil和t的第二个元素
	}

	return nil, token // 否则，返回nil和原始的token值
}
