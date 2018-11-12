package scheduler

import (
	"github.com/zhaojigang/crawler/model"
)

type QueuedScheduler struct {
	requestChann chan model.Request
	// 每一个Worker都有一个自己的chan Request
	// workerChan中存放的是Worker们的chan
	workerChan chan chan model.Request
}

func (s *QueuedScheduler) WorkerChann() chan model.Request {
	return make(chan model.Request)
}

func (s *QueuedScheduler) Submit(request model.Request) {
	s.requestChann <- request
}

func (s *QueuedScheduler) WorkerReady(w chan model.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	// 初始化 requestChann
	s.requestChann = make(chan model.Request)
	// 初始化 workerChan
	s.workerChan = make(chan chan model.Request)

	// 创建一个 goroutine
	// 1. 进行request以及Worker的chan的存储
	// 2. 分发request到worker的chan中
	go func() {
		var requestQ []model.Request
		var workerQ []chan model.Request
		for {
			var activeRequest model.Request
			var activeWorker chan model.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.requestChann:
				// 如果开始requestQ=nil,append之后就是包含一个r元素的切片
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
				// 进行request的分发
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
