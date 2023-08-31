package FriendHandlerFuncs

import (
	"GoRealTimeChat/Internal/Api/GinHandlerFuncs/BaseHandlerFuncs"
	"GoRealTimeChat/Internal/Dao/FriendDao"
	"GoRealTimeChat/Internal/Enums"
	"GoRealTimeChat/Internal/Services/Dispatch"
	"GoRealTimeChat/Internal/Utils"
	"GoRealTimeChat/Packages/Response"
	"github.com/gin-gonic/gin"
)

type FriendHandler struct {
}

var (
	friendDao FriendDao.FriendDAO
)

// GetFriendsList Index 获取好友列表
func (*FriendHandler) GetFriendsList(ginCtx *gin.Context) {
	// 从上下文中获取id
	id := ginCtx.MustGet("id")

	// 调用friendDao的GetFriendLists方法，获取好友列表
	err, lists := friendDao.GetFriendLists(id)

	// 如果发生错误，则返回获取用户列表失败的错误
	if err != nil {
		Response.FailResponse(Enums.ParamError, "获取用户列表失败").ToJson(ginCtx)
		return
	}

	// 返回成功响应，将好友列表转为JSON格式并发送给客户端
	Response.SuccessResponse(lists).ToJson(ginCtx)
	return
}

// ShowFriendInformation Show 获取好友详情
func (*FriendHandler) ShowFriendInformation(ginCtx *gin.Context) {
	// 调用BaseHandlerFuncs的GetPersonId方法，获取person的信息
	err, person := BaseHandlerFuncs.GetPersonId(ginCtx)

	// 如果发生错误，则返回带有错误码和错误信息的失败响应，并将响应转为JSON格式发送给客户端
	if err != nil {
		Response.FailResponse(Enums.ParamError, err.Error()).ToJson(ginCtx)
		return
	}

	// 创建FriendDao的实例friendDao
	var friendDao FriendDao.FriendDAO

	// 调用friendDao的GetFriends方法，传入person.ID作为参数，获取好友列表的结果赋值给err和lists变量
	err, lists := friendDao.GetFriends(person.ID)

	// 如果发生错误，则返回一个成功响应，并将响应转为JSON格式发送给客户端
	if err != nil {
		Response.SuccessResponse().ToJson(ginCtx)
		return
	}

	// 返回一个成功响应，并将lists转为JSON格式发送给客户端
	Response.SuccessResponse(&lists).ToJson(ginCtx)
	return
}

// DeleteFriend Delete 删除好友
func (*FriendHandler) DeleteFriend(ginCtx *gin.Context) {
	// 调用handler的GetPersonId方法，获取person的信息
	err, person := BaseHandlerFuncs.GetPersonId(ginCtx)

	// 如果发生错误，则返回带有错误码和错误信息的失败响应，并将响应转为JSON格式发送给客户端
	if err != nil {
		Response.FailResponse(Enums.ParamError, err.Error()).ToJson(ginCtx)
		return
	}

	// 创建FriendDao的实例friendDao
	var friendDao FriendDao.FriendDAO

	// 调用friendDao的DelFriends方法，传入person.ID和cxt.MustGet("id")作为参数，删除好友，并将结果赋值给errs变量
	errs := friendDao.DelFriends(person.ID, ginCtx.MustGet("id"))

	// 如果errs不为nil，表示删除好友发生错误，返回一个带有错误码和错误信息的失败响应，并将响应转为JSON格式发送给客户端
	if errs != nil {
		Response.FailResponse(Enums.ParamError, errs.Error()).ToJson(ginCtx)
		return
	}

	// 返回一个成功响应，并将响应转为JSON格式发送给客户端
	Response.SuccessResponse().ToJson(ginCtx)
	return
}

// GetUserStatus 获取好友在线状态
func (*FriendHandler) GetUserStatus(ginCtx *gin.Context) {
	// 调用handler的GetPersonId方法，获取person的信息
	err, person := BaseHandlerFuncs.GetPersonId(ginCtx)

	// 如果发生错误，则返回带有错误码和错误信息的失败响应，并将响应转为JSON格式发送给客户端
	if err != nil {
		Response.FailResponse(Enums.ParamError, err.Error()).ToJson(ginCtx)
		return
	}

	var _dispatch Dispatch.DispatchService

	// 调用_dispatch的IsDispatchNode方法，传入person.ID作为参数，判断用户是否为分派节点，结果赋值给ok变量
	ok, _ := _dispatch.IsDispatchNode(person.ID)

	// 如果是分派节点，则返回一个带有用户在线状态和用户ID的成功响应，并将响应转为JSON格式发送给客户端
	if ok {
		Response.SuccessResponse(&UserStatus{
			Status: Enums.WsUserOnline,
			Id:     Utils.StringToInt(person.ID),
		}).ToJson(ginCtx)
		return
	}

	// 如果不是分派节点，则返回一个带有用户离线状态和用户ID的成功响应，并将响应转为JSON格式发送给客户端
	Response.SuccessResponse(&UserStatus{
		Status: Enums.WsUserOffline,
		Id:     Utils.StringToInt(person.ID),
	}).ToJson(ginCtx)
	return
}
