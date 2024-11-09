package main

import (
	"fmt"
	"time"
)

func sleep1AndSayHi() {
	time.Sleep(time.Second)
	fmt.Println("hi")
}

func sleep2AndSayHi() {
	time.Sleep(2 * time.Second)
	fmt.Println("hi")
}

func main() {
	waiter := &WaitGroupWrapper{}
	waiter.Wrap(sleep1AndSayHi)
	waiter.Wrap(sleep2AndSayHi)
	waiter.Wait()
}
