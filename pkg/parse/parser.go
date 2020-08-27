package parse

import (
	"encoding/json"
	"github.com/KamilSwiech/promadg/pkg/tpl"
	"github.com/spf13/viper"
	"os"
	"text/template"
)

func ParseJson(body []byte) {
	rulesPage := JsonToRulesPage(body)
	RulesPageToMD(rulesPage)
}

func JsonToRulesPage(body []byte) RulesPage {
	var rulesPage RulesPage
	json.Unmarshal([]byte(body), &rulesPage)
	return rulesPage
}

func RulesPageToMD(rules RulesPage) {
	filename := viper.GetString("template")
	var tmpl *template.Template
	if filename == "" {
		defaultTpl := tpl.DefaultTemplate()
		tmpl = template.Must(template.New("default").Parse(string(defaultTpl)))
	} else {
		tmpl = template.Must(template.New(filename).ParseFiles(filename))
	}
	err := tmpl.Execute(os.Stdout, rules)
	if err != nil {
		panic(err)
	}
}
