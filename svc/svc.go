package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/*
https://github.com/judwhite/go-svc
包装通用的服务启停逻辑,监听操作系统信号量等逻辑
*/

type Service interface {
	Start() error
	Stop() error
	Context() context.Context
}

func ServiceRun(service Service, sigs ...os.Signal) error {
	if err := service.Start(); err != nil {
		return err
	}
	if len(sigs) == 0 {
		sigs = []os.Signal{syscall.SIGINT, syscall.SIGTERM}
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, sigs...)

	select {
	case _ = <-signalChan:
		fmt.Println("sig")
	case _ = <-service.Context().Done():
	}
	return service.Stop()
}
