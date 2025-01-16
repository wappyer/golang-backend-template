package model

const TableNameAdmin = "admin" // 管理员表

// Admin mapped from table <admin>
type Admin struct {
	Id             string         `gorm:"column:id;type:varchar(20);primaryKey" json:"id"`                                     // 管理员用户id
	Name           string         `gorm:"column:name;type:varchar(255);not null;default:'';comment:管理员用户名" json:"name"`        // 管理员用户名
	Phone          string         `gorm:"column:phone;type:varchar(255);not null;default:'';comment:管理员手机号" json:"phone"`      // 管理员手机号
	Password       string         `gorm:"column:password;type:varchar(255);not null;default:'';comment:管理员密码" json:"password"` // 管理员密码
	TimestampField TimestampField `gorm:"embedded" json:"-"`
}

// TableName Admin's table name
func (*Admin) TableName() string {
	return TableNameAdmin
}
