package errno

// 底层错误 1000 ～ 1999

const (
	CodeSuccess = 0
	CodeUnknown = 1000 + iota
	CodeNetwork
	CodeSystem
)

func InitBaseErrno() {
	RegisterBatch([]Errno{
		{HttpStatus: 200, Code: CodeSuccess, Msg: "成功"},
		{500, CodeUnknown, "未知错误", nil},
		{HttpStatus: 500, Code: CodeNetwork, Msg: "网络异常，请稍后重试"},
		{HttpStatus: 500, Code: CodeSystem, Msg: "系统异常"},
	})
}
