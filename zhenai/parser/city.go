package parser

import (
	"github.com/zhaojigang/crawler/model"
	"regexp"
)

// match[1]=url match[2]=name
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

// 解析单个城市 - 获取单个城市的用户列表
func ParseCity(contents []byte) model.ParseResult {
	result := model.ParseResult{}
	rg := regexp.MustCompile(cityRe)
	allSubmatch := rg.FindAllSubmatch(contents, -1)
	for _, m := range allSubmatch {
		name := string(m[2])
		result.Items = append(result.Items, "user "+name)
		result.Requests = append(result.Requests, model.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) model.ParseResult {
				return ParseProfile(c, name) // 函数式编程，使用函数包裹函数
			},
		})
	}

	return result
}
