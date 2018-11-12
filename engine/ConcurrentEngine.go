package engine

import (
	"github.com/zhaojigang/crawler/fetcher"
	"github.com/zhaojigang/crawler/model"
	"github.com/zhaojigang/crawler/scheduler"
	"log"
)

// 并发引擎
type ConcurrentEngine struct {
	// 调度器
	Scheduler scheduler.Scheduler
	// 开启的 worker 数量
	WorkerCount int
	// item 通道
	ItemChan chan interface{}
}

func (e *ConcurrentEngine) Run(seeds ...model.Request) {
	// 初始化 Scheduler 的队列，并启动配对 goroutine
	e.Scheduler.Run()
	out := make(chan model.ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		// 每个 Worker 都创建自己的一个 chan Request
		createWorker(e.Scheduler.WorkerChann(), out, e.Scheduler);
	}
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out // 阻塞获取
		for _, item := range result.Items {
			log.Printf("ItemSaver getItems, items: %v", item)
			//go func() {
			//	e.ItemChan <- item
			//}()
		}

		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan model.Request, out chan model.ParseResult, notifier scheduler.ReadyNotifier) {
	go func() {
		for {
			notifier.WorkerReady(in)
			r := <-in // 阻塞等待获取
			result, err := worker(r)
			if err != nil {
				continue
			}
			out <- result // 阻塞发送
		}
	}()
}

func worker(r model.Request) (model.ParseResult, error) {
	log.Printf("fetching url:%s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("fetch error, url: %s, err: %v", r.Url, err)
		return model.ParseResult{}, nil
	}
	return r.ParserFunc(body), nil
}
