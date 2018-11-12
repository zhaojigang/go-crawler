package scheduler

import "github.com/zhaojigang/crawler/model"

type ReadyNotifier interface {
	WorkerReady(chan model.Request)
}
