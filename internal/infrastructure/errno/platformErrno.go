package errno

// 平台错误 2000 ～ 2999

const (
	CodeService            = 2000 + iota // 服务异常
	CodeValid                            // 参数校验异常
	CodeDatabase                         // 数据异常
	CodeConfig                           // 业务配置异常
	CodeFile                             // 文件操作异常
	CodeRetData                          // 数据返回异常
	CodeRepeatRequest                    // 重复请求
	CodeIllegalRequest                   // 非法请求
	CodeIllegalRoleRequest               // 非法角色请求

	CodeDataSelect // 数据不存在
	CodeDataInsert // 数据新增失败
	CodeDataDelete // 数据删除失败
	CodeDataUpdate // 数据修改失败
)

func InitPlatformErrno() {
	RegisterBatch([]Errno{
		{HttpStatus: 500, Code: CodeService, Msg: "服务异常"},
		{HttpStatus: 500, Code: CodeValid, Msg: "校验异常"},
		{HttpStatus: 500, Code: CodeDatabase, Msg: "数据异常"},
		{HttpStatus: 500, Code: CodeConfig, Msg: "配置异常"},
		{HttpStatus: 500, Code: CodeFile, Msg: "文件操作异常"},
		{HttpStatus: 500, Code: CodeRetData, Msg: "数据返回异常"},
		{HttpStatus: 500, Code: CodeRepeatRequest, Msg: "请勿重复请求"},
		{HttpStatus: 500, Code: CodeIllegalRequest, Msg: "非法请求"},
		{HttpStatus: 500, Code: CodeIllegalRoleRequest, Msg: "非法角色请求"},
		{HttpStatus: 404, Code: CodeDataSelect, Msg: "数据不存在"},
		{HttpStatus: 500, Code: CodeDataInsert, Msg: "数据新增失败"},
		{HttpStatus: 500, Code: CodeDataDelete, Msg: "数据删除失败"},
		{HttpStatus: 500, Code: CodeDataUpdate, Msg: "数据修改失败"},
	})
}
