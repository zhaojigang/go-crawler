package model

// 请求任务封装体
type Request struct {
	// 需爬取的 Url
	Url string
	// Url 对应的解析函数
	ParserFunc func([]byte) ParseResult
}

// 解析结果
type ParseResult struct {
	// 解析出来的多个 Request 任务
	Requests []Request
	// 解析出来的实体（例如，城市名），是任意类别（interface{}，类比 java Object）
	Items    []interface{}
}
