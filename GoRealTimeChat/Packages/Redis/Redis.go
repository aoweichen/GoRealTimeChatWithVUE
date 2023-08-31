// 包名为Redis
package Redis

// 导入所需包
import (
	"GoRealTimeChat/DataModels/Internal/ConfigModels"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"time"
)

// REDISDB 为全局变量，表示Redis客户端
var REDISDB *redis.Client

// InitClient 函数用于初始化Redis客户端连接
func InitClient() {
	// 创建一个新的Redis客户端实例，并配置相关参数
	REDISDB = redis.NewClient(&redis.Options{
		Network:      "tcp",                                                                         // 使用TCP网络连接
		Addr:         ConfigModels.ConfigData.Redis.Host + ":" + ConfigModels.ConfigData.Redis.Port, // Redis服务器的地址和端口号
		Password:     ConfigModels.ConfigData.Redis.Password,                                        // Redis服务器的密码
		DB:           ConfigModels.ConfigData.Redis.DB,                                              // 要连接的Redis数据库编号
		PoolSize:     ConfigModels.ConfigData.Redis.Poll,                                            // 连接池大小，默认为CPU核心数的四倍
		MinIdleConns: ConfigModels.ConfigData.Redis.Conn,                                            // 在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量
		DialTimeout:  5 * time.Second,                                                               // 连接超时时间
		ReadTimeout:  5 * time.Second,                                                               // 读取超时时间
		WriteTimeout: 5 * time.Second,                                                               // 写入超时时间
		PoolTimeout:  5 * time.Second,                                                               // 连接池超时时间
	})

	// 发送Ping命令到Redis服务器，检查连接是否正常
	_, err := REDISDB.Ping().Result()
	// 如果连接出错，则记录错误日志并退出程序
	if err != nil {
		zap.S().Errorln(err)
	}
}
