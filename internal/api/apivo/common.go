package apivo

type CommonPageReq struct {
	PageSize int `json:"pageSize" form:"pageSize" validate:"required|default:10"` // 每页行数
	Page     int `json:"page" form:"page" validate:"required|default:1"`          // 页码
}
