package apivo

// ExampleLoginReq 后台管理员登录
type ExampleLoginReq struct {
	Phone    string `json:"phone" validate:"required|isMobile"` // 手机号
	Password string `json:"password" validate:"required"`       // 密码
}

type ExampleLoginResp struct {
	Token string `json:"token"`
}
