package model

const TableNameLog = "log"

// Log mapped from table <log>
type Log struct {
	RequestId string  `gorm:"column:request_id;type:varchar(20);primaryKey" json:"requestId"`                       // 请求id
	UserId    string  `gorm:"column:user_id;type:varchar(255);not null;default:'';comment:用户id" json:"userId"`      // 用户id
	Method    string  `gorm:"column:method;type:varchar(255);not null;default:'';comment:请求方法" json:"method"`       // 请求方法
	Route     string  `gorm:"column:route;type:varchar(255);not null;default:'';comment:请求路由" json:"route"`         // 路由
	Path      string  `gorm:"column:path;type:varchar(255);not null;default:'';comment:请求路径" json:"path"`           // 请求路径
	ClientIp  string  `gorm:"column:client_ip;type:varchar(255);not null;default:'';comment:客户端ip" json:"clientIp"` // 客户端ip
	ServerIp  string  `gorm:"column:server_ip;type:varchar(255);not null;default:'';comment:服务端ip" json:"serverIp"` // 服务端ip
	HttpCode  int     `gorm:"column:http_code;type:int(11);not null;default:0;comment:http状态码" json:"httpCode"`     // http状态码
	Code      int     `gorm:"column:code;type:int(11);not null;default:0;comment:系统状态码" json:"code"`                // 系统状态码
	Message   string  `gorm:"column:message;type:varchar(255);not null;default:'';comment:返回描述" json:"message"`     // 返回描述
	Cost      float64 `gorm:"column:cost;type:decimal(12,6);not null;default:0;comment:耗时" json:"cost"`             // 耗时
	ReqTime   string  `gorm:"column:req_time;type:varchar(255);not null;default:'';comment:请求时间" json:"reqTime"`    // 请求时间
	RespTime  string  `gorm:"column:resp_time;type:varchar(255);not null;default:'';comment:响应时间" json:"respTime"`  // 响应时间
}

// TableName Log's table name
func (*Log) TableName() string {
	return TableNameLog
}
