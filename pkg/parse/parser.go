package parse

import (
	"encoding/json"
	"github.com/Masterminds/sprig"
	"github.com/kamilswiec/promadg/pkg/tpl"
	"github.com/spf13/viper"
	"os"
	"path"
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
		name := path.Base(filename)
		tmpl = template.Must(template.New(name).Funcs(fmap).ParseFiles(filename))
	}
	err := tmpl.Execute(os.Stdout, rules)
	if err != nil {
		panic(err)
	}
}
