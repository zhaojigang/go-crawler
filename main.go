package main

import (
	"github.com/zhaojigang/crawler/engine"
	"github.com/zhaojigang/crawler/model"
	"github.com/zhaojigang/crawler/scheduler"
	"github.com/zhaojigang/crawler/zhenai/parser"
)

func main() {
	engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 1000,
	}.Run(model.Request{
		// 种子 Url
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
