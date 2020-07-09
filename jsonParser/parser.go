package jsonParser

import (
	"encoding/json"
)

func JsonToAlertGroups(body []byte) []AlertGroup {
	var obj map[string]interface{}
	json.Unmarshal([]byte(body), &obj)
	alerts := obj["data"].(interface{}).(map[string]interface{})["groups"].([]interface{})
	alertGroups := make([]AlertGroup, len(alerts))
	for k, v := range alerts {
		alertGroups[k] = CreateAlertGroup(v.(map[string]interface{}))

	}
	return alertGroups
}

func JsonToRulesPage(body []byte) RulesPage {
	var rulesPage RulesPage
	json.Unmarshal([]byte(body), &rulesPage)
	return rulesPage
}

func CreateAlertGroup(z map[string]interface{}) AlertGroup {
	name := z["name"].(string)
	file := z["file"].(string)
	rules := ParseRules(z["rules"].([]interface{}))
	interval := z["interval"].(float64)
	return AlertGroup{name, file, rules, interval}
}

func ParseRules(rules []interface{}) []Alert {
	alerts := make([]Alert, len(rules))
	for k, v := range rules {
		rule := v.(map[string]interface{})
		alerts[k] = ExtractAlertFromRule(rule)
	}
	return alerts
}

func ExtractAlertFromRule(rule map[string]interface{}) Alert {
	rawLabels := rule["labels"].(map[string]interface{})
	rawAnnotations := rule["annotations"].(map[string]interface{})
	name := rule["name"].(string)
	query := rule["query"].(string)
	duration := rule["duration"].(float64)
	labels := ParseDescription(rawLabels)
	annotations := ParseDescription(rawAnnotations)
	return Alert{name, labels, annotations, query, duration}
}

func ParseDescription(desc map[string]interface{}) map[string]string {
	m := make(map[string]string)
	for k, v := range desc {
		m[k] = v.(string)
	}
	return m
}
