package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
	"time"
)

// TimestampField
// **********
// 表结构通用时间戳字段
// **********
type TimestampField struct {
	CreateTime time.Time `gorm:"column:create_time;type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;autoUpdateTime" json:"updateTime"`
}

// CommonRespBody
// **********
// 通用结构
// **********
type CommonRespBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type CommonReq struct {
	Id int `json:"id"` // id
}

type CommonListReq struct {
	Id []int `json:"id"`
}

type CommonRes struct {
	Id   int    `json:"id"`
	Name string `json:"name"` // 名称
}

// MyTime
// **********
// 自定义时间类型（主要为了接口json直接返回标准时间格式 2006-01-02 15:04:05）
// **********
type MyTime time.Time

func NewMyTimeNow() MyTime {
	return MyTime(time.Now())
}
func NewMyTimeParseNormal(value string) MyTime {
	return NewMyTimeParse("2006-01-02 15:04:05", value)
}
func NewMyTimeParse(layout, value string) MyTime {
	t, _ := time.Parse(layout, value)
	return MyTime(t)
}

func (t *MyTime) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02 15:04:05", timeStr)
	*t = MyTime(t1)
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02 15:04:05"))
	if formatted == "\"0001-01-01 00:00:00\"" {
		formatted = "\"\""
	}
	return []byte(formatted), nil
}

func (t MyTime) Value() (driver.Value, error) {
	// MyTime 转换成 time.Time 类型
	tTime := time.Time(t)
	return tTime.Format("2006-01-02 15:04:05"), nil
}

func (t *MyTime) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyTime(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyTime) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}

func (t *MyTime) FormatDefault() string {
	return time.Time(*t).Format("2006-01-02 15:04:05")
}

func (t *MyTime) Format(layout string) string {
	return time.Time(*t).Format(layout)
}

func (t *MyTime) NotEmpty() bool {
	begin, _ := time.Parse("2006-01-02 15:04:05", "1900-01-01 00:00:00")
	return time.Time(*t).After(begin)
}

// MyDate
// **********
// 自定义日期类型（主要为了接口json直接返回标准日期格式 2006-01-02）
// **********
type MyDate time.Time

func NewMyDateNow() MyDate {
	return MyDate(time.Now())
}
func NewMyDateParseNormal(value string) MyDate {
	return NewMyDateParse("2006-01-02", value)
}
func NewMyDateParse(layout, value string) MyDate {
	t, _ := time.Parse(layout, value)
	return MyDate(t)
}

func (t *MyDate) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var err error
	//前端接收的时间字符串
	str := string(data)
	//去除接收的str收尾多余的"
	timeStr := strings.Trim(str, "\"")
	t1, err := time.Parse("2006-01-02", timeStr)
	*t = MyDate(t1)
	return err
}

func (t MyDate) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%v\"", time.Time(t).Format("2006-01-02"))
	if formatted == "\"0001-01-01\"" {
		formatted = "\"\""
	}
	return []byte(formatted), nil
}

func (t MyDate) Value() (driver.Value, error) {
	// MyDate 转换成 time.Time 类型
	tDate := time.Time(t)
	return tDate.Format("2006-01-02"), nil
}

func (t *MyDate) Scan(v interface{}) error {
	switch vt := v.(type) {
	case time.Time:
		// 字符串转成 time.Time 类型
		*t = MyDate(vt)
	default:
		return errors.New("类型处理错误")
	}
	return nil
}

func (t *MyDate) String() string {
	return fmt.Sprintf("hhh:%s", time.Time(*t).String())
}

func (t *MyDate) FormatDefault() string {
	return time.Time(*t).Format("2006-01-02")
}

func (t *MyDate) Format(layout string) string {
	return time.Time(*t).Format(layout)
}

func (t *MyDate) NotEmpty() bool {
	begin, _ := time.Parse("2006-01-02", "1900-01-01")
	return time.Time(*t).After(begin)
}
