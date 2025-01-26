package main

import (
	"fmt"
	"log"
	"strconv"
)

/*
https://github.com/kubeedge/kubeedge
在做一些步骤执行逻辑的时候,打印步骤顺序的日志
*/

func NewStep() *Step {
	return &Step{n: 0}
}

type Step struct {
	n int
}

func (s *Step) Printf(format string, args ...interface{}) {
	s.n++
	format = strconv.Itoa(s.n) + ". " + format
	log.Println(fmt.Sprintf(format, args...))
}

func main() {
	step := NewStep()
	step.Printf("one....")
	step.Printf("two....")
	step.Printf("three....")
}
