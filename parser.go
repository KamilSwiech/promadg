package main

import (
	"encoding/json"
)

func JsonToRulesPage(body []byte) RulesPage {
	var rulesPage RulesPage
	json.Unmarshal([]byte(body), &rulesPage)
	return rulesPage
}
