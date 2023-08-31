package MessageHandlerFunc

import (
	"GoRealTimeChat/DataModels/Internal/DataModels/IMGroupMessages"
	"GoRealTimeChat/DataModels/packages/Model"
	"GoRealTimeChat/Internal/Utils"
	"GoRealTimeChat/Packages/Response"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
)

type GroupMessageHandler struct {
}

// sortByGroupMessage函数用于根据发送时间对IMGroupMessages.ImGroupMessages列表进行排序
// 参数:
// - messageList: IMGroupMessages.ImGroupMessages对象列表，需要排序的列表
func (*GroupMessageHandler) sortByGroupMessage(messageList []IMGroupMessages.ImGroupMessages) {
	sort.Slice(messageList, func(i, j int) bool {
		return messageList[i].SendTime < messageList[j].SendTime // 根据发送时间升序排序
	})
}

// GetGroupList GetGroupList函数用于获取群组消息列表
// 参数:
// - ctx: gin.Context对象，用于处理HTTP请求和响应
func (GMH *GroupMessageHandler) GetGroupList(ctx *gin.Context) {
	var groupMessageList []IMGroupMessages.ImGroupMessages // 定义变量groupMessageList，用于存储IMGroupMessages.ImGroupMessages对象列表
	var total int64                                        // 定义变量total，用于存储消息总数

	// 获取页码、群组ID和每页大小
	page, groupID, pageSize := ctx.Query("page"), ctx.Query("to_id"),
		Utils.StringToInt(ctx.DefaultQuery("pageSize", "50"))

	// 构建查询条件，并按发送时间降序排序
	query := Model.MYSQLDB.Model(&IMGroupMessages.ImGroupMessages{}).Preload("Users").
		Where("group_id=?", groupID).Order("send_time desc")

	// 统计匹配的消息总数
	query.Count(&total)

	// 根据页码设置查询条件
	if len(page) > 0 {
		query = query.Where("id<?", page)
	}

	// 查询群组消息列表
	if result := query.Limit(pageSize).Find(&groupMessageList); result.RowsAffected == 0 {
		// 如果查询结果为空，则返回空列表响应
		Response.SuccessResponse(gin.H{
			"list": struct{}{},
			"mate": gin.H{
				"pageSize": pageSize,
				"page":     page,
				"total":    0,
			}}, http.StatusOK).ToJson(ctx)
		return
	}

	// 对群组消息列表进行排序
	GMH.sortByGroupMessage(groupMessageList)

	// 返回群组消息列表响应
	Response.SuccessResponse(gin.H{
		"list": groupMessageList,
		"mate": gin.H{
			"pageSize": pageSize,
			"page":     page,
			"total":    total,
		}}, http.StatusOK).ToJson(ctx)
	return
}
