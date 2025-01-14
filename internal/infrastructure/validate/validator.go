package validate

import (
	"github.com/gookit/validate"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

func addValidator() {
	validate.AddValidators(map[string]interface{}{
		"isIDCard":         validatorIDCard,         // 身份证验证
		"isMobile":         validatorMobile,         // 手机号验证
		"trueName":         validatorTrueName,       // 手机号验证
		"isNormalDate":     validatorNormalDate,     // 日期格式验证
		"isNormalDatetime": validatorNormalDateTime, // 时间格式验证
	})
}

func validatorIDCard(id string) bool {
	// 身份证位数不对
	if len(id) != 15 && len(id) != 18 {
		return false
	}

	// 转大写
	id = strings.ToUpper(id)

	if len(id) == 18 {
		// 验证算法
		if !checkValidNo18(id) {
			return false
		}
	}

	return true
}

// 18位身份证校验码
func checkValidNo18(id string) bool {
	var weight = [17]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	var validValue = [11]byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	id18 := []byte(id)
	nSum := 0
	for i := 0; i < len(id18)-1; i++ {
		n, _ := strconv.Atoi(string(id18[i]))
		nSum += n * weight[i]
	}
	//mod得出18位身份证校验码
	mod := nSum % 11
	if validValue[mod] == id18[17] {
		return true
	}

	return false
}

func validatorMobile(mobile string) bool {
	// 匹配规则
	// ^1第一位为一
	// [345789]{1} 后接一位345789 的数字
	// \\d \d的转义 表示数字 {9} 接9位
	// $ 结束符
	regRuler := "^1[345789]{1}\\d{9}$"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)

	// 返回 MatchString 是否匹配
	return reg.MatchString(mobile)
}

func validatorNormalDate(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	if err == nil {
		return true
	}

	return false
}

func validatorNormalDateTime(datetime string) bool {
	_, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err == nil {
		return true
	}

	return false
}

func validatorTrueName(name string) bool {
	// 需为真实姓名
	// 长度 2～6
	// 仅含有中文
	length := utf8.RuneCountInString(name)
	if length < 2 || length > 6 {
		return false
	}
	for _, i := range name {
		if !unicode.Is(unicode.Han, i) {
			return false
		}
	}
	return true
}
