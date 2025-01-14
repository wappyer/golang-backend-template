package apivo

// AdminLoginReq 后台管理员登录
type AdminLoginReq struct {
	Phone    string `json:"phone" validate:"required|isMobile"` // 手机号
	Password string `json:"password" validate:"required"`       // 密码
}

type AdminLoginResp struct {
	Token string `json:"token"`
}
