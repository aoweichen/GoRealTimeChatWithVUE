package GRPCService

import (
	"GoRealTimeChat/DataModels/Internal/ConfigModels"
	"GoRealTimeChat/Internal/Services/Client/GRPCMessageClient/GRPC/MessageGRPC"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

var GRPCServer = grpc.NewServer()

// StartGRPCServer 用于启动 gRPC 服务器。
func StartGRPCServer() {
	if ConfigModels.ConfigData.Server.ServiceOpen { // 检查配置中的服务开启标志
		var message MessageGRPC.IMGRPCMessage                    // 创建 IMGRPCMessage 实例
		MessageGRPC.RegisterImMessageServer(GRPCServer, message) // 将 IMGRPCMessage 实例注册到 GRPCServer

		listener, err := net.Listen("tcp", ":8002") // 监听指定端口
		if err != nil {
			zap.S().Fatal("grpc服务启动失败！", err) // 如果启动失败，输出错误信息并终止程序
		}

		err = GRPCServer.Serve(listener) // 启动 gRPC 服务器并开始接受连接
		if err != nil {
			return
		}
	}
}
