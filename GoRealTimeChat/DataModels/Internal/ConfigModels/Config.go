package ConfigModels

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Server ServerConfig
	Mysql  MysqlConfig
	Log    LogConfig
	JWT    JWTConfig
	Redis  RedisConfig
	Mail   MailConfig
	Kafka  KafkaConfig
	Nsq    NsqConfig
	QiNiu  QiNiuConfig
	Github GithubConfig
	GoBot  BotConfig
	Gitee  GiteeConfig
}

// ServerConfig ServerConf http server服务的config
// 定义一个结构体 ServerConf，用于存储服务器的配置信息
type ServerConfig struct {
	Name          string `json:"name"`          // 服务器名称，用于标识服务器
	Listen        string `json:"listen"`        // 服务器监听地址，指定服务器监听的 IP 和端口号
	Mode          string `json:"mode"`          // 服务器运行模式，例如开发模式、测试模式等
	Env           string `json:"env"`           // 服务器环境变量，用于配置服务器运行时的环境变量
	Lang          string `json:"lang"`          // 服务器语言设置，指定服务器使用的编程语言或框架
	CoroutinePoll int    `json:"coroutinePoll"` // 协程轮询时间间隔，用于设置协程的轮询间隔时间
	Node          string `json:"node"`          // 节点信息，用于标识服务器所在的物理节点或虚拟节点
	ServiceOpen   bool   `json:"serviceOpen"`   // 集群服务是否开启，用于指示服务器是否正在运行或已停止
	GrpcListen    string `json:"grpcListen"`    // gRPC 服务的监听地址，指定 gRPC 服务监听的 IP 和端口号
	FilePath      string `json:"filePath"`      // 文件路径，用于指定服务器读取配置文件的路径
}

// JWTConfig 是一个结构体，用于存储 JSON Web Token（JWT）的配置信息
type JWTConfig struct {
	Secret     string `json:"secret"`       // Secret 字段表示用于签名和验证的密钥，类型为字符串
	TimeToLive int64  `json:"time_to_live"` // TTL 字段表示 JWT 的有效期，单位为秒，类型为 int64
}

// MysqlConfig MysqlConf 是一个结构体，用于存储 MySQL 数据库的配置信息
type MysqlConfig struct {
	Host     string `json:"host"`     // Host 字段表示 MySQL 数据库的主机地址，类型为字符串
	Port     string `json:"port"`     // Port 字段表示 MySQL 数据库的端口号，类型为字符串
	Username string `json:"username"` // Username 字段表示连接 MySQL 数据库所使用的用户名，类型为字符串
	Password string `json:"password"` // Password 字段表示连接 MySQL 数据库所使用的密码，类型为字符串
	Database string `json:"database"` // Database 字段表示要连接的 MySQL 数据库的名称，类型为字符串
	Charset  string `json:"charset"`  // Charset 字段表示连接 MySQL 数据库时使用的字符集，类型为字符串
}

// LogConfig LogConf 结构体用于存储日志相关的配置信息
type LogConfig struct {
	Level     string `json:"level"`     // 日志级别
	Type      string `json:"type"`      // 日志类型
	FileName  string `json:"filename"`  // 日志文件名
	MaxSize   int    `json:"maxSize"`   // 日志文件最大大小（字节）
	MaxBackup int    `json:"maxBackup"` // 日志文件最大备份数
	MaxAge    int    `json:"maxAge"`    // 日志文件最大保存时间（天）
	Compress  bool   `json:"compress"`  // 是否启用日志压缩
}

// RedisConfig RedisConf 结构体用于存储 Redis 数据库的配置信息
type RedisConfig struct {
	Host     string `json:"host"`     // Redis 服务器主机地址
	Port     string `json:"port"`     // Redis 服务器端口号
	Password string `json:"password"` // Redis 服务器密码
	DB       int    `json:"db"`       // 要连接的 Redis 数据库索引
	Poll     int    `json:"poll"`     // Redis 服务器轮询间隔时间（毫秒）
	Conn     int    `json:"conn"`     // Redis 服务器最大连接数
}

// MailConfig MailConf 结构体用于存储邮件发送相关的配置信息
type MailConfig struct {
	Driver                        string `json:"driver"`                             // 邮件驱动
	Host                          string `json:"host"`                               // 邮件服务器主机地址
	Name                          string `json:"name"`                               // 邮件服务器名称
	Port                          int    `json:"port"`                               // 邮件服务器端口号
	Password                      string `json:"password"`                           // 邮件服务器密码
	Encryption                    string `json:"encryption"`                         // 邮件加密方式
	FromName                      string `json:"fromName"`                           // 发件人名称
	EmailCodeSubject              string `json:"email_code_subject"`                 // 子主题
	EmailCodeHtmlTemplateFilePath string `json:"email_code_html_template_file_path"` // EmailCodeHtml 模板文件地址
}

// KafkaConfig KafkaConf 结构体用于存储 Kafka 消息队列的配置信息
type KafkaConfig struct {
	Host string `json:"host"` // Kafka 服务器主机地址
	Port string `json:"port"` // Kafka 服务器端口号
}

// NsqConfig NsqConf 结构体用于存储 NSQ 消息队列的配置信息
type NsqConfig struct {
	LookupHost string `json:"lookupHost"` // NSQ 服务器主机地址
	NsqHost    string `json:"nsqHost"`    // NSQ 服务器主机地址
}

type QiNiuConfig struct {
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secretKey"`
	Bucket    string `json:"bucket"`
	Domain    string `json:"domain"`
}

type GithubConfig struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	RedirectUrl  string `json:"redirectUrl"`
}

type BotConfig struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
}

type GiteeConfig struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	RedirectUrl  string `json:"redirectUrl"`
}

var ConfigData = &Config{}

// InitConfigs 初始化配置信息
func InitConfigs(configFilePath string) *Config {
	//	设置文件类型为yaml
	viper.SetConfigType("yaml")
	zap.S().Infof("设置文件类型为: yaml")
	//	读取配置文件
	viper.SetConfigFile(configFilePath)
	zap.S().Infof("读取存储在 %#v 路径下的配置文件。", configFilePath)
	// 判断读取配置是否成功读取配置
	viperReadConfigError := viper.ReadInConfig()
	if viperReadConfigError != nil {
		zap.S().Errorf("viper 读取存储在 %#v 的配置文件出错，请检查配置文件是否写错了！", configFilePath)
		panic(viperReadConfigError)
	}
	//	解析配置文件到Struct
	viperUnmarshalError := viper.Unmarshal(&ConfigData)
	if viperUnmarshalError != nil {
		zap.S().Errorln("viper 解析配置文件到Struct出错，请检查配置文件是否写错了！")
		panic(viperUnmarshalError)
	}
	// 配置加载成功
	zap.S().Infof("配置加载成功！")
	//热加载配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		viperUnmarshalError := viper.Unmarshal(&ConfigData)
		if viperUnmarshalError != nil {
			zap.S().Errorln("viper 解析配置文件到Struct出错，请检查配置文件是否写错了！")
			panic(viperUnmarshalError)
		}
	})

	return ConfigData
}

func IsLocal() bool {
	return ConfigData.Log.Level == "local"
}

// IsProduction TODO 准被删除
func IsProduction() bool {
	return ConfigData.Log.Level == "production"
}

// IsTesting TODO 准被删除
func IsTesting() bool {
	return ConfigData.Log.Level == "testing"
}
