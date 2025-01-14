package errno

import (
	"errors"
	"fmt"
	"sync"
)

func Initialize() {
	InitBaseErrno()
	InitPlatformErrno()
	InitBusinessErrno()
}

// Errno 统一错误err结构
type Errno struct {
	HttpStatus int
	Code       int
	Msg        string
	Err        error
}

func (c Errno) GetHttpStatus() int {
	return c.HttpStatus
}

func (c Errno) GetCode() int {
	return c.Code
}

func (c Errno) GetMsg() string {
	return c.Msg
}

func (c Errno) Error() string {
	return c.Msg
}

func (c Errno) NotNil() bool {
	return c.Code != 0
}

func (c Errno) IsNil() bool {
	return c.Code == 0
}

var codes = map[int]Errno{} // 初始化时存储所有注册的错误码
var codesMux sync.Mutex

// Register 注册错误码
func Register(httpStatus, code int, msg string) {
	codesMux.Lock()
	defer codesMux.Unlock()
	codes[code] = Errno{httpStatus, code, msg, nil}
}

// RegisterBatch 批量注册错误码
func RegisterBatch(errs []Errno) {
	codesMux.Lock()
	defer codesMux.Unlock()
	for _, e := range errs {
		codes[e.Code] = e
	}
}

func NewErrno(code int) Errno {
	if Errno, ok := codes[code]; ok {
		return Errno
	}
	return Errno{500, CodeUnknown, "未知错误", nil}
}

func NewErrnoWithMsg(code int, format string, args ...interface{}) Errno {
	errno := NewErrno(code)
	if len(args) > 0 {
		errno.Err = fmt.Errorf(format, args)
	} else {
		errno.Err = errors.New(format)
	}
	errno.Msg = errno.Err.Error()
	return errno
}

func NewErrnoWithErr(code int, err error) Errno {
	errno := NewErrno(code)
	if err != nil {
		errno.Err = err
		errno.Msg = errno.Err.Error()
	}
	return errno
}
