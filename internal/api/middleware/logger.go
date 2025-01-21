package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"gitee.com/wappyer/golang-backend-template/config"
	"gitee.com/wappyer/golang-backend-template/global"
	"gitee.com/wappyer/golang-backend-template/internal/api/controller"
	"gitee.com/wappyer/golang-backend-template/internal/domain/entity"
	model "gitee.com/wappyer/golang-backend-template/internal/infrastructure/db/model"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/db/repository"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/logger"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/uid"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func Logger(serverConf config.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		var requestId = uid.LongUid(1).GetStr()
		c.Set(global.ContextKeyTraceId, requestId)

		// 打印路由日志
		reqMethod := c.Request.Method
		reqURI := c.Request.RequestURI
		reqBody, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			zap.S().Error(err)
		}
		if len(reqBody) == 0 {
			reqBody = []byte("{}")
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 请求记录到log表
		serverIp := strconv.Itoa(serverConf.Index) + "-" + serverConf.Port
		reqLog := ""
		if c.Request.URL.RawQuery != "" {
			reqLog = c.Request.URL.RawQuery
		} else if len(reqBody) > 5000 {
			reqLog = "data too long"
		} else {
			reqLog = string(reqBody)
		}

		err = new(repository.LogRepository).Insert(c, &entity.Log{
			Log:    &model.Log{RequestId: requestId, Method: reqMethod, Route: c.FullPath(), Path: c.Request.URL.Path, ClientIp: c.ClientIP(), ServerIp: serverIp, ReqTime: startTime.Format("2006-01-02 15:04:05.000000")},
			Detail: &model.LogDetail{RequestId: requestId, Req: reqLog},
		})

		c.Next()

		statusCode := c.Writer.Status()
		latencyTime := float64(time.Since(startTime)) / 1000 / 1000

		handlerInfo := &HandlerInfo{
			Id:          requestId,
			LatencyTime: latencyTime,
			Request: Request{
				Header: RequestHeader{
					Method:      reqMethod,
					RequestURI:  reqURI,
					ContentType: c.ContentType(),
					ClientIP:    c.ClientIP(),
					UserAgent:   c.Request.UserAgent(),
				},
				Body: JsonRawString(reqBody),
			},
			Response: Response{
				Header: ResponseHeader{
					Status: statusCode,
				},
				Body: JsonRawString(blw.body.Bytes()),
			},
		}
		if strings.Contains(c.Writer.Header().Get("Content-Type"), "application/json") {
			data, _ := json.Marshal(handlerInfo)
			logger.InfoF(c, "[request] %s", data)
		} else {
			logger.InfoF(c, "[request] %s", handlerInfo.Request.Header)
		}

		// 更新log表响应
		respCode, respMessage, respSaveFlag := getRespCodeAndMessage(c, blw.body.Bytes(), c.FullPath())
		resp := blw.body.String()
		if !respSaveFlag {
			resp = resp[:200]
		}

		go (new(repository.LogRepository)).Update(c, &entity.Log{
			Log:    &model.Log{RequestId: requestId, UserId: c.GetString(global.ContextKeyLoginId), HttpCode: c.Writer.Status(), Code: respCode, Message: respMessage, Cost: latencyTime, RespTime: time.Now().Format("2006-01-02 15:04:05.000000")},
			Detail: &model.LogDetail{RequestId: requestId, Resp: resp},
		})
	}
}

// HandlerInfo 自定义gin路由日志
type HandlerInfo struct {
	Id          string   `json:"id"`
	LatencyTime float64  `json:"latency_time"`
	Request     Request  `json:"request"`
	Response    Response `json:"response"`
}
type Request struct {
	Header RequestHeader `json:"header"`
	Body   JsonRawString `json:"body"`
}
type Response struct {
	Header ResponseHeader `json:"header"`
	Body   JsonRawString  `json:"body"`
}

type RequestHeader struct {
	Method          string `json:"method"`
	RequestURI      string `json:"request_uri"`
	ContentType     string `json:"content_type"`
	ClientIP        string `json:"client_ip"`
	UserAgent       string `json:"user_agent"`
	RequestUserName string `json:"request_user_name"`
	RequestUserId   int    `json:"request_user_id"`
}
type ResponseHeader struct {
	Status int `json:"status"`
}

type JsonRawString string

func (m JsonRawString) MarshalJSON() ([]byte, error) {
	if len(m) == 0 {
		return []byte("null"), nil
	}
	return []byte(m), nil
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// 获取接口返回数据中的code，message
func getRespCodeAndMessage(ctx context.Context, body []byte, route string) (int, string, bool) {
	code, message, respSaveFlag := 0, "", true
	if string(body) == "" {
		return code, message, true
	}

	// 本地接口返回
	resp := &controller.RespBody{}
	if err := json.Unmarshal(body, resp); err != nil {
		message = "非标准结构返回数据"
		respSaveFlag = false
		return code, message, respSaveFlag
	}
	code = resp.Code
	message = resp.Message
	return code, message, respSaveFlag
}
