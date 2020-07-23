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
	viper.WriteConfig()
}
