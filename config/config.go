package config

var Conf Config

// Config 组合全部配置模型
type Config struct {
	Env     Env
	Server  Server     `mapstructure:"server"`
	Monitor Server     `mapstructure:"monitor"`
	Docs    Server     `mapstructure:"docs"`
	Db      DB         `mapstructure:"db"`
	Redis   Redis      `mapstructure:"redis"`
	Log     LogConfigs `mapstructure:"log"`
	Wechat  Wechat     `mapstructure:"wechat"`
	Oss     Oss        `mapstructure:"oss"`
	Sms     Sms        `mapstructure:"sms"`
	Jwt     Jwt        `mapstructure:"jwt"`
	MQ      MQ         `mapstructure:"mq"`
}

// Server 服务启动端口号配置
type Server struct {
	Enable bool   `mapstructure:"enable"`
	Host   string `mapstructure:"host"`
	Port   string `mapstructure:"port"`
	Index  int    `mapstructure:"index"`
}

// DB MySQL数据源配置
type DB struct {
	DbType      string `mapstructure:"dbType"`      // 数据库类型 mysql
	Path        string `mapstructure:"path"`        // 服务器地址
	Dbname      string `mapstructure:"dbName"`      // 数据库名
	Config      string `mapstructure:"config"`      // 高级配置
	Username    string `mapstructure:"username"`    // 登录名
	Password    string `mapstructure:"password"`    // 密码
	MaxIdleConn int    `mapstructure:"maxIdleConn"` // 空闲中的最大连接数
	MaxOpenConn int    `mapstructure:"maxOpenConn"` // 打开到数据库的最大连接数
	RedisEnable bool   `mapstructure:"redisEnable"` // 是否开启redis
}

// Redis 缓存
type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Db       int    `mapstructure:"db"`
}

// LogConfigs 日志相关
type LogConfigs struct {
	LogLevel          string `mapstructure:"logLevel"`          // 日志打印级别 debug  info  warning  error
	LogFormat         string `mapstructure:"logFormat"`         // 输出日志格式	logfmt, json
	LogPath           string `mapstructure:"logPath"`           // 输出日志文件路径
	LogFileName       string `mapstructure:"logFileName"`       // 输出日志文件名称
	LogFileMaxSize    int    `mapstructure:"logFileMaxSize"`    // 【日志分割】单个日志文件最多存储量 单位(mb)
	LogFileMaxBackups int    `mapstructure:"logFileMaxBackups"` // 【日志分割】日志备份文件最多数量
	LogMaxAge         int    `mapstructure:"logMaxAge"`         // 日志保留时间，单位: 天 (day)
	LogCompress       bool   `mapstructure:"logCompress"`       // 是否压缩日志
	LogStdout         bool   `mapstructure:"logStdout"`         // 是否输出到控制台
}

// Wechat 微信小程序相关配置
type Wechat struct {
	AppId           string `mapstructure:"appId"`
	AppSecret       string `mapstructure:"appSecret"`
	MchId           string `mapstructure:"mchid"`           // 商户ID
	SerialNo        string `mapstructure:"serialNo"`        // 商户证书的证书序列号
	ApiV3Key        string `mapstructure:"apiV3Key"`        // APIv3Key,商户平台获取
	PrivateKey      string `mapstructure:"privateKey"`      // 商户API证书下载后，私钥 apiclient_key.pem 读取后的字符串内容
	PayNotifyUrl    string `mapstructure:"payNotifyUrl"`    // 微信支付结果通知回调url
	RefundNotifyUrl string `mapstructure:"refundNotifyUrl"` // 微信退款结果通知回调url
	DoctorAppId     string `mapstructure:"doctorAppId"`     // 医生小程序appid
	DoctorAppSecret string `mapstructure:"doctorAppSecret"` // 医生小程序appsecret
	Redis           Redis  `mapstructure:"redis"`           // 缓存配置
}

// Oss oss相关
type Oss struct {
	EasyCheck OssConf `mapstructure:"easyCheck"`
	Lis       OssConf `mapstructure:"lis"`
}

type OssConf struct {
	Bucket           string `mapstructure:"bucket"`
	Endpoint         string `mapstructure:"endpoint"`
	EndpointInternal string `mapstructure:"endpointInternal"`
	AccessKeyId      string `mapstructure:"accessKeyId"`
	AccessKeySecret  string `mapstructure:"accessKeySecret"`
	CallbackUrl      string `mapstructure:"callbackUrl"`
	ExpireTime       int64  `mapstructure:"expireTime"`
}

type Sms struct {
	ApiKey string `mapstructure:"apiKey"`
}

type Jwt struct {
	SigningKey string `mapstructure:"signingKey"`
}

type MQ struct {
	Nats Nats `mapstructure:"nats"`
}

type Nats struct {
	Url               string `mapstructure:"url"`
	MaxMsgsPerSubject int64  `mapstructure:"maxMsgsPerSubject"` // 每个主题最大消息数
}
