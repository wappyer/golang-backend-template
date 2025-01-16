package errno

// 平台错误 3000 ～ 9999

const (
	CodeCreateToken            = 3000 + iota // 初始化token失败
	CodeRequestTokenNil                      // 用户未登录
	CodeIllegalToken                         // token异常，身份认证失败
	CodeLoginTimeout                         // 登录超时
	CodeUserInfo                             // 用户信息异常
	CodeWechatCode                           // 微信身份验证异常
	CodeSmsValidCodeDomain                   // 非法请求短信验证码
	CodeFilePutToOss                         // 文件上传失败
	CodeSmsSend                              // 发送短信失败
	CodeNoValidCode                          // 无效验证码
	CodeSmsValidCodeStatusUsed               // 验证码已使用
	CodeSmsValidCodeExpired                  // 验证码已过期
)

func InitBusinessErrno() {
	RegisterBatch([]Errno{
		{HttpStatus: 500, Code: CodeCreateToken, Msg: "初始化token失败"},
		{HttpStatus: 401, Code: CodeRequestTokenNil, Msg: "用户未登录"},
		{HttpStatus: 401, Code: CodeIllegalToken, Msg: "token异常，身份认证失败"},
		{HttpStatus: 401, Code: CodeLoginTimeout, Msg: "登录超时"},
		{HttpStatus: 401, Code: CodeUserInfo, Msg: "用户信息异常"},
		{HttpStatus: 401, Code: CodeWechatCode, Msg: "微信身份验证异常"},
		{HttpStatus: 500, Code: CodeSmsValidCodeDomain, Msg: "非法请求短信验证码"},
		{HttpStatus: 500, Code: CodeFilePutToOss, Msg: "文件上传失败"},
		{HttpStatus: 500, Code: CodeSmsSend, Msg: "发送短信失败"},
		{HttpStatus: 404, Code: CodeNoValidCode, Msg: "无效验证码"},
		{HttpStatus: 500, Code: CodeSmsValidCodeStatusUsed, Msg: "验证码已使用，请重新发送"},
		{HttpStatus: 500, Code: CodeSmsValidCodeExpired, Msg: "验证码已过期，请重新发送"},
	})
}
