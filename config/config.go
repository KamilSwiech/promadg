package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var ViperConfig Config

func InitConfiguration(v *viper.Viper) error {
	v.SetConfigName("config")         // name of config file (without extension)
	v.SetConfigType("yaml")           // REQUIRED if the config file does not have the extension in the name
	v.AddConfigPath("/etc/promadg/")  // path to look for the config file in
	v.AddConfigPath("$HOME/.promadg") // call multiple times to add many search paths
	v.AddConfigPath(".")              // optionally look for config in the working directory
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			return fmt.Errorf("Config File not Found")
		} else {
			return fmt.Errorf("Yaml configuration file probably spoiled")
		}
	}
	ViperConfig.Configuration = v
	return nil
}

type Context struct {
	Prometheus string
}

type Config struct {
	Configuration *viper.Viper
}

func LoadContext(v *viper.Viper) (*Context, error) {
	if val := ViperConfig.Configuration.GetString("prometheus"); val != "" {
		return &Context{val}, nil
	}
	return nil, fmt.Errorf("No context in config file")
}

func GetConfig() (v Config) {
	return ViperConfig
}
