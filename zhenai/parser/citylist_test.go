package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	expectRequestsLen := 470
	expectCitiesLen := 470
	// 表格驱动测试
	expectRequestUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectRequestCities := []string{
		"city 阿坝",
		"city 阿克苏",
		"city 阿拉善盟",
	}

	body, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(body)

	if len(result.Requests) != expectRequestsLen {
		t.Errorf("expect requestLen %d, but %d", expectRequestsLen, len(result.Requests))
	}
	if len(result.Items) != expectCitiesLen {
		t.Errorf("expect citiesLen %d, but %d", expectCitiesLen, len(result.Items))
	}

	for i, url := range expectRequestUrls {
		if url != result.Requests[i].Url {
			t.Errorf("expect url %s, but %s", url, result.Requests[i].Url)
		}
	}

	for i, city := range expectRequestCities {
		if city != result.Items[i] {
			t.Errorf("expect url %s, but %s", city, result.Items[i])
		}
	}
}
