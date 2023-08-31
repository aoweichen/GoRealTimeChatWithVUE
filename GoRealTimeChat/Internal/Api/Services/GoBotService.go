package Services

import (
	"GoRealTimeChat/DataModels/Internal/ApiRequests"
	"GoRealTimeChat/DataModels/Internal/ConfigModels"
	"GoRealTimeChat/DataModels/Internal/DataModels/IMUser"
	"GoRealTimeChat/DataModels/packages/Model"
	"GoRealTimeChat/Internal/Dao/MessageDao"
	"GoRealTimeChat/Internal/Enums"
	"GoRealTimeChat/Internal/Utils"
	"GoRealTimeChat/Packages/Date"
	"GoRealTimeChat/Packages/Hash"
	"fmt"
	"strings"
	"sync"
)

var (
	messagesServices ImMessageService
	botData          = map[string]string{} // 存储指令
	lock             sync.RWMutex
	BOT_NOE          = 1
	messageDao       MessageDao.MessageDao
)

// InitChatBot 初始化聊天机器人
func InitChatBot() {
	var count int64

	// 查询指定ID的用户数量
	Model.MYSQLDB.Table("im_users").Where("id=?", BOT_NOE).Count(&count)

	// 如果数量为0，表示用户不存在，则创建用户
	if count == 0 {
		hashedPassword, _ := Hash.SaltCryptoHashPassword(ConfigModels.ConfigData.GoBot.Password)
		createdAt := Date.NewDate()

		// 创建用户记录
		Model.MYSQLDB.Table("im_users").Create(&IMUser.ImUsers{
			ID:            int64(BOT_NOE),
			Email:         ConfigModels.ConfigData.GoBot.Email,
			Password:      hashedPassword,
			Name:          ConfigModels.ConfigData.GoBot.Name,
			CreatedAt:     createdAt,
			UpdatedAt:     createdAt,
			Avatar:        ConfigModels.ConfigData.GoBot.Avatar,
			LastLoginTime: createdAt,
			Uid:           Utils.GetUuid(),
			UserJson:      "{}",
			UserType:      1,
		})
	}
}

// GetMessage 获取指定关键字的消息
func GetMessage(key string) string {

	// 如果关键字包含冒号
	if strings.Contains(key, ":") {
		arr := strings.Split(key, ":")

		// 如果切割后的数组长度为2，表示格式正确
		if len(arr) == 2 {
			lock.Lock()
			botData[arr[0]] = arr[1]
			lock.Unlock()
			return "很不错就是这样~"
		}

		// 如果切割后的数组长度大于2，表示格式不正确
		if len(arr) > 2 {
			return "格式不对呀~"
		}
	}

	// 如果关键字在botData中存在，返回对应的值
	if value, ok := botData[key]; ok {
		return value
	} else {
		// 如果关键字在botData中不存在，返回默认提示语
		return "没明白您的意思-暂时还不知道说啥~~~ 你可以通过 xxx:xxx 指令定义消息😊"
	}
}

// InitChatBotMessage 初始化聊天机器人的消息
func InitChatBotMessage(fromID int64, toID int64) {
	params := ApiRequests.PrivateMessageRequest{
		MsgId:       Date.TimeUnixNano(),
		MsgCode:     Enums.WsChatMessage,
		MsgClientId: Date.TimeUnixNano(),
		FromID:      fromID,
		ToID:        toID,
		ChannelType: 1,
		MsgType:     1,
		Message:     fmt.Sprintf("您好呀~ 我是%s~🥰", ConfigModels.ConfigData.GoBot.Name),
		SendTime:    Date.NewDate(),
		Data:        "",
	}

	// 创建消息记录
	messageDao.CreateMessage(params)

	// 发送私聊消息
	messagesServices.SendPrivateMessage(params)

	// 设置新的消息内容
	params.Message = "我们来玩个游戏吧！你问我答~！👋"

	// 创建消息记录
	messageDao.CreateMessage(params)

	// 发送私聊消息
	messagesServices.SendPrivateMessage(params)
}
