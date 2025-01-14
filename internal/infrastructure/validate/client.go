package validate

import (
	"encoding/json"
	"github.com/gookit/validate"
	"github.com/gookit/validate/locales/zhcn"
	"reflect"
)

func Initialize() {
	// 加载参数校验语言包
	zhcn.RegisterGlobal()

	// 手动添加全局消息
	changeMessage()

	// 添加自定义验证器
	addValidator()
}

// Bind 验证gin中请求参数并绑定至obj
func Bind[T any](obj T, check func(arg T) error, rs ...any) (err error) {
	if va := validate.Struct(obj); !va.Validate() {
		return va.Errors.OneError()
	}

	if check != nil {
		if e := check(obj); e != nil {
			return
		}
	}

	if len(rs) > 0 {
		var last any = obj
		for i := range rs {
			b, e := json.Marshal(last)
			if e != nil {
				return e
			}
			err = json.Unmarshal(b, &rs[i])
			if err != nil {
				return
			}
			last = rs[i]
		}
	}

	// 处理分页
	if _, ok := reflect.TypeOf(obj).Elem().FieldByName("Page"); ok {
		objValue := reflect.ValueOf(obj).Elem()
		page := objValue.FieldByName("Page").Int()
		pageSize := objValue.FieldByName("PageSize").Int()
		offset := pageSize * (page - 1)
		limit := pageSize
		for i := range rs {
			reflect.ValueOf(rs[i]).Elem().FieldByName("Offset").SetInt(offset)
			reflect.ValueOf(rs[i]).Elem().FieldByName("Limit").SetInt(limit)
		}
	}

	return nil
}
