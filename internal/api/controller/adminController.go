package controller

import (
	"gitee.com/wappyer/golang-backend-template/internal/api/apivo"
	"gitee.com/wappyer/golang-backend-template/internal/application"
	"gitee.com/wappyer/golang-backend-template/internal/application/appvo"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	AdminApp *application.AdminApplication
}

func NewAdminController() *AdminController {
	return &AdminController{
		AdminApp: application.NewAdminApplication(),
	}
}

func (c *AdminController) Login(ctx *gin.Context) {
	param := appvo.LoginReq{}
	if err := Bind(ctx, &apivo.AdminLoginReq{}, nil, &param); err != nil {
		FailedByValid(ctx, err)
		return
	}

	data, eno := c.AdminApp.Admin(ctx.Request.Context(), param)
	if eno.NotNil() {
		Failed(ctx, eno)
		return
	}

	FormatAndReturn(ctx, &apivo.AdminLoginResp{}, data)
	return
}
