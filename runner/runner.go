package runner

import (
	"context"
	"time"
)

type RunnersFactory struct {
	Runners []Runner
}

type Runner interface {
	Initialize() error
	Run()
	Shutdown(context.Context)
}

var runnersFactory *RunnersFactory

func NewRunnersFactory() *RunnersFactory {
	runnersFactory = &RunnersFactory{
		Runners: make([]Runner, 0),
	}
	return runnersFactory
}

func (r *RunnersFactory) RegisterRunner(runner Runner) *RunnersFactory {
	if err := runner.Initialize(); err == nil {
		r.Runners = append(r.Runners, runner)
	}
	return runnersFactory
}

func (r *RunnersFactory) Run() {
	for _, runner := range r.Runners {
		go runner.Run()
	}
}

func (r *RunnersFactory) Shutdown() {
	for _, runner := range r.Runners {
		ctx, cancelApi := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelApi()
		runner.Shutdown(ctx)
	}
}
