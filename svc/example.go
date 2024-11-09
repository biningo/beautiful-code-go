package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type HttpServer struct {
	addr       string
	httpServer *http.Server
}

func (s *HttpServer) Start() error {
	fmt.Println("start........")
	server := &http.Server{Addr: s.addr, Handler: nil}
	s.httpServer = server

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
	return nil
}

func (s *HttpServer) Stop() error {
	fmt.Println("stopping...")
	return s.httpServer.Shutdown(context.Background())
}

func (server *HttpServer) Context() context.Context {
	return context.Background()
}

func main() {
	httpService := &HttpServer{addr: ":8080"}
	if err := ServiceRun(httpService); err != nil {
		log.Println(err)
	}
}
