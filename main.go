package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func main() {
	initConfiguration()
	// Subcommands
	if err := InitCmds(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
	viper.WriteConfig()
}
