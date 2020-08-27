package main

import (
	"encoding/json"
	"github.com/spf13/viper"
	"os"
	"text/template"
)

func JsonToRulesPage(body []byte) RulesPage {
	var rulesPage RulesPage
	json.Unmarshal([]byte(body), &rulesPage)
	return rulesPage
}

func RulesPageToMD(rules RulesPage) {
	filename := viper.GetString("template")
	if filename == "" {

	}
	tmpl, err := template.New(filename).ParseFiles(filename)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, rules)
	if err != nil {
		panic(err)
	}
}
