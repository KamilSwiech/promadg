package main

import (
	"fmt"
	"github.com/spf13/viper"
	"kamilswiech/promadg/config"
	"kamilswiech/promadg/jsonParser"
	"kamilswiech/promadg/markdown"
	"kamilswiech/promadg/parseCmd"
	"kamilswiech/promadg/promApi"
	"os"
)

func main() {
	v := viper.New()
	config.InitConfiguration(v)
	z, err := config.LoadContext(v)
	fmt.Println(z.Prometheus)
	// Subcommands
	if err := cmd.InitCmds(os.Args, v); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	z, err = config.LoadContext(v)
	fmt.Println(z.Prometheus)

	url := "http://ksa-prometheus.blue/api/v1/rules?type=alert"
	resp, err := promapi.SendRequest(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := promapi.GetJson(resp)
	if err != nil {
		return
	}
	alertGroups := jsonParser.JsonToAlertGroups(body)
	_ = markdown.AlertGroupsToMarkdown(alertGroups)
	//fmt.Println(markdown)
	v.WriteConfig()
}
