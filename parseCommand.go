package main

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func InitCmds(args []string, v *viper.Viper) error {
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	setCmd := flag.NewFlagSet("set", flag.ExitOnError)

	setCtxPtr := setCmd.String("context", "", "Prometheus instance context as url")
	setConfPtr := setCmd.String("config", "", "Configuration for printing requests")
	setTimePtr := setCmd.String("time", "", "Time interval for requesting data")

	getDataPtr := getCmd.String("data", "", "Prometheus data to scrape")
	getTimePtr := getCmd.String("time", "", "Get time interval which is used for scraping data")
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

	var setFunc = func(data string, v *viper.Viper) {}
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
		setFunc(data, v)
	}
	if getCmd.Parsed() {
		if *getDataPtr != "" {
			data = *getDataPtr
			getFunc = GetData
		} else if *getTimePtr != "" {
			data = *getTimePtr
			getFunc = GetTime
		} else if *getURLPtr != "" {
			data = *getURLPtr
			getFunc = GetURL
		}
		getFunc(data)
	}
	return nil
}

func SetCtx(data string, v *viper.Viper) {
	fmt.Println("SetCtx")
	v.Set("prometheus", data)
}

func SetConf(data string, v *viper.Viper) {
	fmt.Println("SetConf")
	v.Set("configuration", data)
}

func SetTime(data string, v *viper.Viper) {
	fmt.Println("SetTime")
	v.Set("time", data)
}

func GetData(data string) {
	fmt.Println("GetData")
}

func GetTime(data string) {
	fmt.Println("GetTime")

}

func GetURL(data string) {
	fmt.Println("GetURL")
}
