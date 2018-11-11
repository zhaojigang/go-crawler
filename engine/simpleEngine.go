package engine

import (
	"github.com/zhaojigang/crawler/fetcher"
	"github.com/zhaojigang/crawler/model"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...model.Request) {
	// Request 任务队列
	var requests []model.Request
	// 将 seeds Request 放入 []requests，即初始化 []requests
	for _, r := range seeds {
		requests = append(requests, r)
	}
	// 执行任务
	for len(requests) > 0 {
		// 1. 获取第一个 Request，并从 []requests 移除，实现了一个队列功能
		r := requests[0]
		requests = requests[1:]

		parseResult, err := e.worker(r)
		if err != nil {
			continue
		}
		// 4. 将解析体中的 []Requests 加到请求任务队列 requests 的尾部
		requests = append(requests, parseResult.Requests...)

		// 5. 遍历解析出来的实体，直接打印
		for _, item := range parseResult.Items {
			log.Printf("getItems, url: %s, items: %v", r.Url, item)
		}
	}
}

func (e SimpleEngine) worker(r model.Request) (model.ParseResult, error) {
	// 2. 使用爬取器进行对 Request.Url 进行爬取
	body, err := fetcher.Fetch(r.Url)
	// 如果爬取出错，记录日志
	if err != nil {
		log.Printf("fetch error, url: %s, err: %v", r.Url, err)
		return model.ParseResult{}, nil
	}

	// 3. 使用 Request 的解析函数对怕渠道的内容进行解析
	return r.ParserFunc(body), nil
}
