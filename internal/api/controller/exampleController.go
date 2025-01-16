package controller

import (
	"gitee.com/wappyer/golang-backend-template/internal/api/apivo"
	"gitee.com/wappyer/golang-backend-template/internal/application"
	"gitee.com/wappyer/golang-backend-template/internal/application/appvo"
	"github.com/gin-gonic/gin"
)

type ExampleController struct {
	ExampleApp *application.ExampleApplication
}

func NewExampleController() *ExampleController {
	return &ExampleController{
		ExampleApp: application.NewExampleApplication(),
	}
}

func (c *ExampleController) Example(ctx *gin.Context) {
	param := appvo.ExampleReq{}
	if err := Bind(ctx, &apivo.ExampleLoginReq{}, nil, &param); err != nil {
		FailedByValid(ctx, err)
		return
	}

	data, eno := c.ExampleApp.Example(ctx.Request.Context(), param)
	if eno.NotNil() {
		Failed(ctx, eno)
		return
	}

	FormatAndReturn(ctx, &apivo.ExampleLoginResp{}, data)
	return
}
