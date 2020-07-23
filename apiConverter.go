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
	var str strings.Builder
	prometheus := viper.GetString("prometheus")
	str.WriteString(Protocol)
	str.WriteString(prometheus)
	str.WriteString(ApiPrefix)
	var query string
	if data == "rules" {
		query = "rules"
	} else if data == "alerts" {
		query = "rules?type=alert"
	}
	str.WriteString(query)
	return str.String()
}
