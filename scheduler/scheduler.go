package scheduler

import (
	"github.com/zhaojigang/crawler/model"
)

// 调度器接口
type Scheduler interface {
	ReadyNotifier
	Submit(request model.Request)
	WorkerChann() chan model.Request
	Run()
}
