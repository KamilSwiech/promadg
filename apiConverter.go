package main

import (
	"github.com/spf13/viper"
	"strings"
)

const (
	Protocol  = "http://"
	ApiPrefix = "/api/v1/"
)

func ParseRequest(data string) string {
	var query string
	var markdown string
	if data == "rules" {
		query = "rules"
		body := ExecuteUrl(query)
		GetRules(body)
	} else if data == "alerts" {
		query = "rules?type=alert"
		body := ExecuteUrl(query)
		GetRules(body)
	}
	return markdown
}

func BuildUrl(data string) string {
	var str strings.Builder
	prometheus := viper.GetString("prometheus")
	str.WriteString(Protocol)
	str.WriteString(prometheus)
	str.WriteString(ApiPrefix)
	str.WriteString(data)
	return str.String()
}

func ExecuteUrl(data string) []byte {
	url := BuildUrl(data)
	return GetRequest(url)
}

func GetRules(body []byte) {
	rules := JsonToRulesPage(body)
	RulesPageToMD(rules)
}
