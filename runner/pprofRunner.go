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
	Env    config.Env
	Conf   config.Server
	Server *http.Server
}

func NewPprofRunner() *PprofRunner {
	return &PprofRunner{}
}

func (r *PprofRunner) Initialize() error {
	if !config.Conf.Monitor.Enable {
		logger.InfoF(context.Background(), "[init] pprof服务未开启")
		return errors.New("pprof服务未开启")
	}
	r.Env = config.Conf.Env
	r.Conf = config.Conf.Monitor
	return nil
}

func (r *PprofRunner) Run() {
	// 启动、监听日志端口
	r.WebServer()
}

func (r *PprofRunner) WebServer() {
	pprofServer := &http.Server{}
	if r.Env != config.Pro {
		pprofGinEngine := gin.New()
		addr := fmt.Sprintf(":%s", r.Conf.Port)
		pprofServer = &http.Server{
			Addr:           addr,
			Handler:        pprofGinEngine,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   40 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		logger.InfoF(context.Background(), "[init] 启动pprof服务 listening at %v", addr)
		err := pprofServer.ListenAndServe()
		if err != nil {
			panic(fmt.Sprintf("[init] 启动pprof服务失败: %s \n", err))
		}
		r.Server = pprofServer
	}
}

func (r *PprofRunner) Shutdown(ctx context.Context) {
	if r.Server == nil {
		return
	}
	if err := r.Server.Shutdown(ctx); err != nil {
		panic(fmt.Sprintf("[init] 退出pprof服务失败: %s \n", err))
	}
}
