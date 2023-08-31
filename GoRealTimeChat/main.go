package main

import (
	"GoRealTimeChat/DataModels/Internal/ConfigModels"
	"GoRealTimeChat/DataModels/packages/Model"
	"GoRealTimeChat/Internal/Api/Services"
	"GoRealTimeChat/Internal/Middleware"
	"GoRealTimeChat/Internal/Router"
	"GoRealTimeChat/Internal/Services/Client/ClientManager"
	"GoRealTimeChat/Internal/Services/GRPCService"
	"GoRealTimeChat/Packages/CoroutinesPoll"
	"GoRealTimeChat/Packages/Logger"
	"GoRealTimeChat/Packages/MessageQueue/NSQQueue"
	"GoRealTimeChat/Packages/Redis"
	"github.com/gin-gonic/gin"
)

func StartService(router *gin.Engine) {
	go ClientManager.IMMessageClientManager.StartServer()    // 在后台启动 IMMessageClientManager 的服务器
	router.Use(Middleware.Recover)                           // 使用 Recover 中间件，用于处理请求发生的 panic
	SetRoute(router)                                         // 设置路由
	gin.SetMode(ConfigModels.ConfigData.Server.Mode)         // 设置 Gin 的运行模式
	GRPCService.StartGRPCServer()                            // 启动 gRPC 服务器
	err := router.Run(ConfigModels.ConfigData.Server.Listen) // 运行 Gin 引擎，监听指定地址和端口
	if err != nil {
		panic(err) // 如果运行出错，则抛出异常
	}
}

func init() {
	ConfigModels.InitConfigs("H:\\PROJECT\\goSTU\\GoRealTimeChat\\Configs\\configs.yaml") // 初始化配置，从指定的配置文件中加载配置
	Logger.InitLogger()                                                                   // 初始化日志记录器
	Model.InitMySQLDB()                                                                   // 初始化 MySQL 数据库连接
	Redis.InitClient()                                                                    // 初始化 Redis 客户端
	CoroutinesPoll.ConnectPool()                                                          // 连接协程池
	err := NSQQueue.InitNewNSQProducerPoll()                                              // 初始化 NSQ 生产者连接池
	if err != nil {
		panic(err) // 如果初始化失败，则抛出异常
	}
	Services.InitChatBot()
}

func SetRoute(router *gin.Engine) {
	Router.RegisterApiRoutes(router) // 注册 API 路由
	Router.RegisterWSRouters(router) // 注册 WebSocket 路由
}

func main() {
	r := gin.Default() // 创建一个默认的 Gin 引擎实例
	StartService(r)    // 启动服务
}
