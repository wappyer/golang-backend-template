package controller

import (
	"gitee.com/wappyer/golang-backend-template/internal/api/apivo"
	"gitee.com/wappyer/golang-backend-template/internal/application"
	"gitee.com/wappyer/golang-backend-template/internal/application/appvo"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	LoginApp *application.LoginApplication
}

func NewAdminController() *AdminController {
	return &AdminController{
		LoginApp: application.NewLoginApplication(),
	}
}

func (c *AdminController) Login(ctx *gin.Context) {
	param := appvo.LoginReq{}
	if err := Bind(ctx, &apivo.AdminLoginReq{}, nil, &param); err != nil {
		FailedByValid(ctx, err)
		return
	}

	data, eno := c.LoginApp.Login(ctx.Request.Context(), param)
	if eno.NotNil() {
		Failed(ctx, eno)
		return
	}

	FormatAndReturn(ctx, &apivo.AdminLoginResp{}, data)
	return
}
