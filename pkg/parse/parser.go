package parse

import (
	"encoding/json"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/kamilswiec/promadg/pkg/tpl"
	"github.com/spf13/viper"
	"os"
	"path"
	"text/template"
)

func ParseJson(body []byte) error {
	rulesPage, err := JsonToRulesPage(body)
	if err != nil {
		return fmt.Errorf("Failed to parse json: %s", err)
	}
	return RulesPageToMD(rulesPage)
}

func JsonToRulesPage(body []byte) (RulesPage, error) {
	var rulesPage RulesPage
	err := json.Unmarshal([]byte(body), &rulesPage)
	return rulesPage, err
}

func RulesPageToMD(rules RulesPage) error {
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
	return tmpl.Execute(os.Stdout, rules)
}
