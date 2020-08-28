package parse

import (
	"encoding/json"
	"github.com/kamilswiec/promadg/pkg/tpl"
	"github.com/spf13/viper"
	"github.com/Masterminds/sprig"
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
	fmap := sprig.TxtFuncMap()
	if filename == "" {
		defaultTpl := string(tpl.DefaultTemplate())
		tmpl = template.Must(
			template.New("default").Funcs(fmap).Parse(defaultTpl))
	} else {
		tmpl = template.Must(template.New(filename).Funcs(fmap).ParseFiles(filename))
	}
	err := tmpl.Execute(os.Stdout, rules)
	if err != nil {
		panic(err)
	}
}
