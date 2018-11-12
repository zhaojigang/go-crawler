package scheduler

import (
	"github.com/zhaojigang/crawler/model"
)

type SimpleScheduler struct {
	workerChan chan model.Request
}

func (s *SimpleScheduler) WorkerChann() chan model.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan model.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan model.Request)
}

func (s *SimpleScheduler) Submit(request model.Request) {
	go func() { s.workerChan <- request }()
}
