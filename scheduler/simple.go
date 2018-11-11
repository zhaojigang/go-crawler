package scheduler

import (
	"github.com/zhaojigang/crawler/model"
)

type SimpleScheduler struct {
	workerChan chan model.Request
}

// 为什么使用指针接收者，需要改变 SimpleScheduler 内部的 workerChan
// https://stackoverflow.com/questions/27775376/value-receiver-vs-pointer-receiver-in-golang
// https://studygolang.com/articles/1113
// https://blog.csdn.net/suiban7403/article/details/78899671
func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan model.Request) {
	s.workerChan = in
}

func (s *SimpleScheduler) Submit(request model.Request) {
	// 每个 Request 一个 Goroutine
	go func() { s.workerChan <- request }()
}
