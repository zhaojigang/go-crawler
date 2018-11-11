package parser

import (
	"github.com/zhaojigang/crawler/model"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]*)</a>`

// cityList 的 ParserFunc func([]byte) ParseResult
// 解析种子页面 - 获取城市列表
func ParseCityList(contents []byte) model.ParseResult {
	result := model.ParseResult{}
	// 正则表达式：()用于提取
	rg := regexp.MustCompile(cityListRe)
	allSubmatch := rg.FindAllSubmatch(contents, -1)
	// 遍历每一个城市的匹配字段（城市 Url 和城市名），并且将 Url 和城市解析器封装为一个 Request
	// 最后将该 Request 添加到 ParseResult 中
	for _, m := range allSubmatch {
		result.Items = append(result.Items, "city "+string(m[2]))
		result.Requests = append(result.Requests, model.Request{
			Url:        string(m[1]),
			ParserFunc: ParseCity,
		})
	}
	// 返回 ParseResult
	return result
}
