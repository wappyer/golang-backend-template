package validate

import "github.com/gookit/validate"

func changeMessage() {
	validate.AddGlobalMessages(map[string]string{
		"isIDCard":         "身份证号校验有误",
		"isMobile":         "手机号校验有误",
		"trueName":         "真实姓名应为2至6个中文",
		"isNormalDate":     "日期格式应为：2006-01-02",
		"isNormalDatetime": "时间格式应为：2006-01-02 15:04:05",
	})
}
