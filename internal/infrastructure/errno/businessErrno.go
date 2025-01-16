package errno

// 平台错误 3000 ～ 9999

const (
	CodeCreateToken                = 3000 + iota // 初始化token失败
	CodeRequestTokenNil                          // 用户未登录
	CodeIllegalToken                             // token异常，身份认证失败
	CodeLoginTimeout                             // 登录超时
	CodeUserInfo                                 // 用户信息异常
	CodeAdminLoginPassword                       // 登录密码错误
	CodeWechatCode                               // 微信身份验证异常
	CodeMember                                   // 无效成员信息
	CodeAddress                                  // 无效地址信息
	CodeUserPhone                                // 请先授权手机号
	CodeUserPhoneDuplicate                       // 手机号已存在
	CodeIdentityVerify                           // 实名认证失败
	CodeInspectionRecord                         // 无效检测报告
	CodeSmsValidCodeDomain                       // 非法请求短信验证码
	CodeFilePutToOss                             // 文件上传失败
	CodeNoInspectionRecord                       // 未查询到检测报告
	CodeUnOneselfInspectionRecord                // 非本人检测报告
	CodeUnPersonalInspectionRecord               // 非手动添加检测报告不可删除
	CodeSmsSend                                  // 发送短信失败
	CodeNoValidCode                              // 无效验证码
	CodeSmsValidCodeStatusUsed                   // 验证码已使用
	CodeSmsValidCodeExpired                      // 验证码已过期
	CodeDoctorPhoneAuthed                        // 手机号已认证
	CodeDoctorPhoneStop                          // 手机号受限制
	CodeDoctorPhoneUnAdd                         // 未录入手机号
)

func InitBusinessErrno() {
	RegisterBatch([]Errno{
		{HttpStatus: 500, Code: CodeCreateToken, Msg: "初始化token失败"},
		{HttpStatus: 401, Code: CodeRequestTokenNil, Msg: "用户未登录"},
		{HttpStatus: 401, Code: CodeIllegalToken, Msg: "token异常，身份认证失败"},
		{HttpStatus: 401, Code: CodeLoginTimeout, Msg: "登录超时"},
		{HttpStatus: 401, Code: CodeUserInfo, Msg: "用户信息异常"},
		{HttpStatus: 401, Code: CodeAdminLoginPassword, Msg: "登录密码错误"},
		{HttpStatus: 401, Code: CodeWechatCode, Msg: "微信身份验证异常"},
		{HttpStatus: 404, Code: CodeMember, Msg: "无效成员信息"},
		{HttpStatus: 404, Code: CodeAddress, Msg: "无效地址信息"},
		{HttpStatus: 500, Code: CodeUserPhone, Msg: "请先授权手机号"},
		{HttpStatus: 500, Code: CodeUserPhoneDuplicate, Msg: "手机号已存在，请联系客服"},
		{HttpStatus: 500, Code: CodeIdentityVerify, Msg: "实名认证失败"},
		{HttpStatus: 404, Code: CodeInspectionRecord, Msg: "无效检测报告"},
		{HttpStatus: 500, Code: CodeSmsValidCodeDomain, Msg: "非法请求短信验证码"},
		{HttpStatus: 500, Code: CodeFilePutToOss, Msg: "文件上传失败"},
		{HttpStatus: 404, Code: CodeNoInspectionRecord, Msg: "未查询到检测报告"},
		{HttpStatus: 404, Code: CodeUnOneselfInspectionRecord, Msg: "非本人检测报告"},
		{HttpStatus: 500, Code: CodeUnPersonalInspectionRecord, Msg: "非手动添加检测报告不可删除"},
		{HttpStatus: 500, Code: CodeSmsSend, Msg: "发送短信失败"},
		{HttpStatus: 404, Code: CodeNoValidCode, Msg: "无效验证码"},
		{HttpStatus: 500, Code: CodeSmsValidCodeStatusUsed, Msg: "验证码已使用，请重新发送"},
		{HttpStatus: 500, Code: CodeSmsValidCodeExpired, Msg: "验证码已过期，请重新发送"},
		{HttpStatus: 500, Code: CodeDoctorPhoneAuthed, Msg: "手机号已认证"},
		{HttpStatus: 500, Code: CodeDoctorPhoneStop, Msg: "手机号受限制"},
		{HttpStatus: 500, Code: CodeDoctorPhoneUnAdd, Msg: "未录入手机号"},
	})
}
