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
}

func (e ConcurrentEngine) Run(seeds ...model.Request) {
	in := make(chan model.Request)
	out := make(chan model.ParseResult)
	// 初始化调度器的 chann
	e.Scheduler.ConfigureMasterWorkerChan(in)
	// 创建 WorkerCount 个 worker
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out);
	}
	// 将 seeds 中的 Request 添加到调度器 chann
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out // 阻塞获取
		for _, item := range result.Items {
			log.Printf("getItems, items: %v", item)
		}

		for _, r := range result.Requests {
			// 如果 submit 内部直接是 s.workerChan <- request，则阻塞等待发送，该方法阻塞在这里
			// 如果 submit 内部直接是 go func() { s.workerChan <- request }()，则为每个Request分配了一个Goroutine，这里不会阻塞在这里
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan model.Request, out chan model.ParseResult) {
	go func() {
		for {
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
