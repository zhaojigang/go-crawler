package main

import (
	"github.com/zhaojigang/crawler/engine"
	"github.com/zhaojigang/crawler/model"
	"github.com/zhaojigang/crawler/persist"
	"github.com/zhaojigang/crawler/scheduler"
	"github.com/zhaojigang/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(model.Request{
		// 种子 Url
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
