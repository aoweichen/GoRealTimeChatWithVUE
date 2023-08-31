package Router

import (
	"GoRealTimeChat/Internal/Api/GinHandlerFuncs/AuthHandlerFuncs"
	"GoRealTimeChat/Internal/Api/GinHandlerFuncs/FriendHandlerFuncs"
	"GoRealTimeChat/Internal/Api/GinHandlerFuncs/GroupHandlerFuncs"
	"GoRealTimeChat/Internal/Api/GinHandlerFuncs/MessageHandlerFunc"
	"GoRealTimeChat/Internal/Api/GinHandlerFuncs/SessionHandlerFuncs"
	"GoRealTimeChat/Internal/Api/GinHandlerFuncs/UploadFileHandlerFunc"
	"GoRealTimeChat/Internal/Api/GinHandlerFuncs/UserHandlerFuncs"
	"GoRealTimeChat/Internal/Middleware"
	"github.com/gin-gonic/gin"
)

var (
	login         AuthHandlerFuncs.AuthHandler
	users         UserHandlerFuncs.UserHandler
	sessions      SessionHandlerFuncs.SessionHandler
	friends       FriendHandlerFuncs.FriendHandler
	friendRecord  FriendHandlerFuncs.FriendRecordHandler
	messagesH     MessageHandlerFunc.MessageHandler
	groupMessageH MessageHandlerFunc.GroupMessageHandler
	groupsH       GroupHandlerFuncs.GroupHandler
	clouds        UploadFileHandlerFunc.QINIUYUNHandler
)

// RegisterApiRoutes 注册api路由
func RegisterApiRoutes(routerEngine *gin.Engine) {

	// 设置允许跨域资源访问
	routerEngine.Use(Middleware.CORS())
	//
	apiGroup := routerEngine.Group("/api/v1")
	{
		//	注册登录相关Group
		authGroup := apiGroup.Group("/auth")
		{
			// 登录接口
			authGroup.POST("/login", login.Login)
			// 验证码发送接口
			authGroup.POST("/sendEmailCode", login.SendEmailCode)
			// 注册账号接口
			authGroup.POST("/registered", login.Registered)
		}
		apiGroup.Use(Middleware.Auth())
		// 用户接口
		{
			// 获得 id 所对应的用户信息
			apiGroup.GET("/user/:id", users.UserInfo)
			// 获得用户对应的通讯录列表，包含好友和群聊
			apiGroup.Any("/contact/list", users.UserContactList)
		}
		// 会话接口
		{
			// 获得某用户所有会话列表
			apiGroup.GET("/sessions/get_list", sessions.GetSessionsList)
			// 新增会话
			apiGroup.POST("/sessions/add", sessions.AddSessions)
			// 更新会话
			apiGroup.PUT("/sessions/update/:id", sessions.UpdateSessions)
			// 删除会话
			apiGroup.DELETE("/sessions/delete/:id", sessions.DeleteSessions)
		}
		// 好友接口
		{
			// 获取相关信息
			apiGroup.Any("/friends/list", friends.GetFriendsList)
			apiGroup.GET("/friends/information/:id", friends.ShowFriendInformation)
			apiGroup.DELETE("/friends/delete/:id", friends.DeleteFriend)
			apiGroup.GET("/friends/status/:id", friends.GetUserStatus)
			// 相关请求
			apiGroup.POST("/friends/add/request", friendRecord.SendAddFriendsRequest)
			apiGroup.GET("/friends/add/request/list", friendRecord.GetFriendsRequestRecord)
			apiGroup.PUT("/friends/add/agree", friendRecord.AgreeOrRejectFriendRequest)
			apiGroup.GET("/friends/query", friendRecord.UserQuery)
		}
		// 消息接口
		{
			apiGroup.GET("/messages/private/list", messagesH.GetPrivateChatList)
			apiGroup.GET("/messages/group/list", groupMessageH.GetGroupList)
			apiGroup.POST("/messages/private", messagesH.SendMessage)
			apiGroup.POST("/messages/group", messagesH.SendMessage)
			apiGroup.POST("/messages/video", messagesH.SendVideoMessage)
			apiGroup.POST("/messages/recall", messagesH.RecallMessage)
		}
		// 群聊接口
		{
			apiGroup.POST("/groups/createGroup", groupsH.CreateGroup)
			apiGroup.POST("/groups/applyJoin/:id", groupsH.ApplyJoin)
			apiGroup.POST("/groups/AddOrRemoveUser", groupsH.AddOrRemoveUser)
			apiGroup.GET("/groups/lists", groupsH.GetGroupList)
			apiGroup.GET("/groups/users/:id", groupsH.GetUsers)
			apiGroup.DELETE("/groups/logout/:id", groupsH.Logout)
		}
		// 文件上传接口
		{
			// TODO 文件上传
			apiGroup.POST("/upload/file", clouds.UploadFile).Use(Middleware.Auth())
		}

	}

}
