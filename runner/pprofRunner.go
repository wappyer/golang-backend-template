package runner

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/wappyer/golang-backend-template/config"
	"gitee.com/wappyer/golang-backend-template/internal/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type PprofRunner struct {
	Conf   config.Server
	Server *http.Server
}

func NewPprofRunner(conf config.Server) *PprofRunner {
	return &PprofRunner{
		Conf: conf,
	}
}

func (r *PprofRunner) Initialize() error {
	if !r.Conf.Enable {
		logger.InfoF(context.Background(), "[init] pprof服务未开启")
		return errors.New("pprof服务未开启")
	}
	return nil
}

func (r *PprofRunner) Run() {
	// 启动、监听日志端口
	r.WebServer()
}

func (r *PprofRunner) WebServer() {
	r.Server = &http.Server{}
	if !config.EnvIsPro() {
		pprofGinEngine := gin.New()
		addr := fmt.Sprintf(":%s", r.Conf.Port)
		r.Server = &http.Server{
			Addr:           addr,
			Handler:        pprofGinEngine,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   40 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		logger.InfoF(context.Background(), "[init] 启动pprof服务 listening at %v", addr)
		err := r.Server.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				logger.InfoF(context.Background(), "[init] pprof服务退出.")
			} else {
				logger.ErrorF(context.Background(), "[init] pprof服务异常中断: %s.", err)
			}
		}
	}
	return
}

func (r *PprofRunner) Shutdown(ctx context.Context) {
	if r.Server == nil {
		return
	}
	if err := r.Server.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("[init] 退出pprof服务失败: %s \n", err))
	}
}
