package parse

import (
	"concurrentCrawler/engine"
	"concurrentCrawler/model"
	"log"
	"regexp"
	"strings"

	"github.com/bitly/go-simplejson"
)

var re = regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

// Profile TODO
func Profile(contents []byte, name string) engine.ParseResult {
	match := re.FindSubmatch(contents)
	result := engine.ParseResult{}
	if len(match) >= 2 {
		json := match[1]
		profile := parseJson(json)
		profile.Name = name
		result.Items = append(result.Items, profile)
	}

	return result

}

// 解析json数据
func parseJson(json []byte) model.Profile {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("解析json失败。。")
	}

	infos, err := res.Get("objectInfo").Get("basicInfo").Array() // 判断是否是数组

	var profile model.Profile
	for k, v := range infos {

		if e, ok := v.(string); ok {
			switch k {
			case 0:
				profile.Marriage = e
			case 1:
				profile.Age = e
			case 2:
				profile.Xingzuo = e
			case 3:
				profile.Height = e
			case 4:
				profile.Weight = e
			case 6:
				profile.Income = e
			case 7:
				profile.Occupation = e
			case 8:
				profile.Education = e
			}
		}

	}

	infos2, err := res.Get("objectInfo").Get("detailInfo").Array()

	for _, v := range infos2 {

		if e, ok := v.(string); ok {
			if strings.Contains(e, "族") {
				profile.Hukou = e
			} else if strings.Contains(e, "房") {
				profile.House = e
			} else if strings.Contains(e, "车") {
				profile.Car = e
			}
		}
	}
	gender, err := res.Get("objectInfo").Get("genderString").String()
	profile.Gender = gender

	return profile

}
