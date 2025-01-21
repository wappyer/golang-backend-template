package db

// Config 数据源配置
type Config struct {
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
