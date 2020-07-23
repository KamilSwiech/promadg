package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitCmds(args []string) error {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	setCmd := flag.NewFlagSet("set", flag.ExitOnError)

	setCtxPtr := setCmd.String("context", "", "Prometheus instance context as url")
	setConfPtr := setCmd.String("config", "", "Configuration for printing requests")
	setTimePtr := setCmd.String("time", "", "Time interval for requesting data")

	getDataPtr := getCmd.String("data", "", "Prometheus data to scrape")
	getTimePtr := getCmd.Bool("time", false, "Get time interval which is used for scraping data")
	getURLPtr := getCmd.String("url", "", "User url to request")

	if len(args) < 2 {
		return fmt.Errorf("Not enough arguments, expected %d, got %d", 2, len(args))
		os.Exit(1)
	}

	switch args[1] {
	case "get":
		getCmd.Parse(args[2:])
	case "set":
		setCmd.Parse(args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	var setFunc = func(data string) {}
	var getFunc = func(data string) {}
	var data string
	if setCmd.Parsed() {
		if *setCtxPtr != "" {
			data = *setCtxPtr
			setFunc = SetCtx
		} else if *setConfPtr != "" {
			data = *setConfPtr
			setFunc = SetConf
		} else if *setTimePtr != "" {
			data = *setTimePtr
			setFunc = SetTime
		}
		setFunc(data)
	}
	if getCmd.Parsed() {
		if *getDataPtr != "" {
			data = *getDataPtr
			getFunc = GetData
		} else if *getTimePtr != false {
			data = ""
			getFunc = GetTime
		} else if *getURLPtr != "" {
			data = *getURLPtr
			getFunc = GetURL
		}
		getFunc(data)
	}
	return nil
}

func SetCtx(data string) {
	fmt.Println("SetCtx")
	viper.Set("prometheus", data)
}

func SetConf(data string) {
	fmt.Println("SetConf")
	viper.Set("configuration", data)
}

func SetTime(data string) {
	fmt.Println("SetTime")
	viper.Set("time", data)
}

func GetData(data string) {
	fmt.Println("GetData")
}

func GetTime(data string) {
	fmt.Println("GetTime")
	fmt.Println(viper.GetString("time"))
}

func GetURL(data string) {
	fmt.Println("GetURL")
	response, err := SendRequest(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	body, err := GetJson(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
