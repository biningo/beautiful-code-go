package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

/**
nsq https://github.com/nsqio/nsq
*/

func demo1() error {
	time.Sleep(time.Second)
	return errors.New("demo1 error")
}

func demo2() error {
	fmt.Println("demo2")
	return nil
}

/*
统一处理error
*/
func main() {
	exitCh := make(chan error)
	var once sync.Once
	exitFunc := func(err error) {
		once.Do(func() {
			if err != nil {
				log.Println("error:" + err.Error())
				exitCh <- err
			}
		})
	}

	go exitFunc(demo1())
	go exitFunc(demo2())

	err := <-exitCh
	fmt.Println("done! error:" + err.Error())
}
