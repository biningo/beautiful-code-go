package main

import "sync"

/**
nsq https://github.com/nsqio/nsq
*/

type WaitGroupWrapper struct {
	sync.WaitGroup
}

/*
包装WaitGroup 让代码更简洁
*/
func (w *WaitGroupWrapper) Wrap(callback func()) {
	w.Add(1)
	go func() {
		callback()
		w.Done()
	}()
}
