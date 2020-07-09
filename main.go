package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	v := viper.New()
	initConfiguration(v)
	z, err := LoadContext(v)
	fmt.Println(z.Prometheus)
	// Subcommands
	if err := InitCmds(os.Args, v); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	z, err = LoadContext(v)
	fmt.Println(z.Prometheus)

	url := "http://ksa-prometheus.blue/api/v1/rules?type=alert"
	resp, err := SendRequest(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := GetJson(resp)
	if err != nil {
		return
	}
	alertGroups := JsonToAlertGroups(body)
	_ = AlertGroupsToMarkdown(alertGroups)
	//fmt.Println(markdown)
	v.WriteConfig()
}
