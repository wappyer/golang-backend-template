package model

const TableNameLogDetail = "log_detail"

// LogDetail mapped from table <log_detail>
type LogDetail struct {
	RequestId      string         `gorm:"column:request_id;type:varchar(20);primaryKey" json:"requestId"` // 请求id
	Req            string         `gorm:"column:req;type:text;not null;comment:请求内容" json:"req"`          // 请求内容
	Resp           string         `gorm:"column:resp;type:text;not null;comment:请求响应" json:"resp"`        // 请求响应
	TimestampField TimestampField `gorm:"embedded" json:"-"`
}

// TableName LogDetail's table name
func (*LogDetail) TableName() string {
	return TableNameLogDetail
}
