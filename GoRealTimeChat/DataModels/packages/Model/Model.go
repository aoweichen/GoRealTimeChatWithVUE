package Model

import (
	"GoRealTimeChat/DataModels/Internal/ConfigModels"
	"GoRealTimeChat/DataModels/Internal/DataModels/IMFriends"
	"GoRealTimeChat/DataModels/Internal/DataModels/IMFriendsRecords"
	"GoRealTimeChat/DataModels/Internal/DataModels/IMGroupMessages"
	"GoRealTimeChat/DataModels/Internal/DataModels/IMGroupUsers"
	"GoRealTimeChat/DataModels/Internal/DataModels/IMGroups"
	"GoRealTimeChat/DataModels/Internal/DataModels/IMMessages"
	"GoRealTimeChat/DataModels/Internal/DataModels/IMOfflineMessage"
	"GoRealTimeChat/DataModels/Internal/DataModels/IMSessions"
	"GoRealTimeChat/DataModels/Internal/DataModels/IMUser"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var MYSQLDB *gorm.DB

// InitMySQLDB 初始化 mysql 数据库
func InitMySQLDB() *gorm.DB {
	//	得到dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		ConfigModels.ConfigData.Mysql.Username,
		ConfigModels.ConfigData.Mysql.Password,
		ConfigModels.ConfigData.Mysql.Host,
		ConfigModels.ConfigData.Mysql.Port,
		ConfigModels.ConfigData.Mysql.Database,
		ConfigModels.ConfigData.Mysql.Charset)
	var gormMySQLConnectedError error
	MYSQLDB, gormMySQLConnectedError = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if gormMySQLConnectedError != nil {
		zap.S().Errorf("Mysql 连接异常: ", gormMySQLConnectedError)
		panic(gormMySQLConnectedError)
	}
	zap.S().Info("mysql 连接成功")

	// 迁移表 IMUser.ImUsers
	MYSQLAutoMigrateDataStruct(&IMUser.ImUsers{})
	// 迁移表 IMSessions.ImSessions{}
	MYSQLAutoMigrateDataStruct(&IMSessions.ImSessions{})
	// 迁移表 IMGroupUsers.ImGroupUsers{}
	MYSQLAutoMigrateDataStruct(&IMGroupUsers.ImGroupUsers{})
	// 迁移表 IMGroups.ImGroups{}
	MYSQLAutoMigrateDataStruct(&IMGroups.ImGroups{})

	// 迁移表 IMOfflineMessage.ImOfflineMessages{}
	MYSQLAutoMigrateDataStruct(&IMOfflineMessage.ImOfflineMessages{})
	// 迁移表 IMOfflineMessage.ImGroupOfflineMessages{}
	MYSQLAutoMigrateDataStruct(&IMOfflineMessage.ImGroupOfflineMessages{})
	// 迁移表 IMMessages.ImMessages{}
	MYSQLAutoMigrateDataStruct(&IMMessages.ImMessages{})
	// 迁移表 IMGroupMessages.ImGroupMessages{}
	MYSQLAutoMigrateDataStruct(&IMGroupMessages.ImGroupMessages{})
	// 迁移表 IMFriendsRecords.ImFriendRecords{}
	MYSQLAutoMigrateDataStruct(&IMFriendsRecords.ImFriendRecords{})
	// 迁移表 IMFriends.ImFriends{}
	MYSQLAutoMigrateDataStruct(&IMFriends.ImFriends{})

	return MYSQLDB

}

// MYSQLAutoMigrateDataStruct
// 定义名为MYSQLAutoMigrateDataStruct的函数
// 该函数用于自动迁移数据结构到MySQL数据库
// 参数dataStruct是一个接口类型，表示要迁移的数据结构
func MYSQLAutoMigrateDataStruct(dataStruct interface{}) {
	// 调用MYSQLDB的AutoMigrate方法，将数据结构迁移到数据库中
	err := MYSQLDB.AutoMigrate(dataStruct)

	// 如果迁移过程中出现错误，则记录错误日志并抛出panic
	if err != nil {
		zap.S().Errorf("迁移表 %#v 出错: %#v", dataStruct, err)
		panic(err)
	}

	// 记录迁移成功的日志
	zap.S().Infof("迁移表 %#v 成功", dataStruct)
}
