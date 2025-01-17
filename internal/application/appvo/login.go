package appvo

import "gitee.com/wappyer/golang-backend-template/internal/domain/entity"

type LoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string       `json:"token"`
	Admin entity.Admin `json:"admin"`
}
