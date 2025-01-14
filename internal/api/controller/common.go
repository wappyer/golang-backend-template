package controller

import (
	"errors"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/errno"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/validate"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

/*******************
 * 输入参数绑定及校验 *
 *******************/

func Bind[T any](ctx *gin.Context, obj T, check func(arg T) error, rs ...any) error {
	// 路由参数绑定
	if err := ctx.ShouldBindUri(obj); err != nil {
		return err
	}

	// 消息体参数绑定
	if err := ctx.ShouldBind(obj); err != nil {
		return err
	}

	// 参数验证
	if err := validate.Bind(obj, check, rs...); err != nil {
		return err
	}
	return nil
}

/******************
 * 格式化数据并输出 *
 ******************/

// FormatAndReturn 格式化返回数据（service层丢出来的数据会是比较全的表数据，接口返回的数据都应该定义返回结构体，只返回需要的字段）
func FormatAndReturn(ctx *gin.Context, toValue interface{}, fromValue interface{}) {
	if err := copier.CopyWithOption(toValue, fromValue, copier.Option{DeepCopy: true}); err != nil {
		Failed(ctx, errno.NewErrno(errno.CodeRetData))
		return
	}
	Success(ctx, toValue)
	return
}

/**************
 * 返回结构定义 *
 **************/

type RespBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	writeResponse(ctx, nil, data)
	return
}

func Failed(ctx *gin.Context, err error) {
	writeResponse(ctx, err, nil)
	return
}

func FailedByCode(ctx *gin.Context, code int) {
	writeResponse(ctx, errno.NewErrno(code), nil)
	return
}

func FailedByValid(ctx *gin.Context, err error) {
	writeResponse(ctx, errno.NewErrnoWithErr(errno.CodeValid, err), nil)
	return
}

func writeResponse(ctx *gin.Context, err error, data interface{}) {
	e := errno.Errno{}

	if err == nil {
		e = errno.NewErrno(errno.CodeSuccess)
	} else if ok := errors.As(err, &e); ok {

	} else {
		e = errno.NewErrnoWithErr(errno.CodeNetwork, err)
	}

	ctx.JSON(e.GetHttpStatus(), RespBody{
		Code:    e.GetCode(),
		Message: e.GetMsg(),
		Data:    data,
	})
	return
}
