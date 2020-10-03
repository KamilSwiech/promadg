package promhandler

import (
	"github.com/spf13/viper"
	"strings"
)

const (
	Protocol  = "http://"
	ApiPrefix = "/api/v1"
	RulesPath = "/rules"
	AlertType = "?type=alert"
)

func BuildRulesUrl() string {
	var str strings.Builder
	prometheus := viper.GetString("prometheus")
	str.WriteString(Protocol)
	str.WriteString(prometheus)
	str.WriteString(ApiPrefix)
	str.WriteString(RulesPath)
	str.WriteString(AlertType)
	return str.String()
}
