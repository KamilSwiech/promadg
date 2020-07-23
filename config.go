package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func initConfiguration() error {
	viper.SetConfigName("config")         // name of config file (without extension)
	viper.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/etc/promadg/")  // path to look for the config file in
	viper.AddConfigPath("$HOME/.promadg") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return fmt.Errorf("Config File not Found")
		} else {
			return fmt.Errorf("Yaml configuration file probably spoiled")
		}
	}
	return nil
}
