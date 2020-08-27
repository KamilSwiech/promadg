package cmd

import (
	"fmt"
	"github.com/kamilswiec/promadg/pkg/parse"
	"github.com/kamilswiec/promadg/pkg/promhandler"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "promadg",
		Short: "Promadg is prometheus alert documentation generator",
		Long: `Promadg allows to create customizable
                documents from prometheus alerts/rules pages.`,
		Run: func(cmd *cobra.Command, args []string) {
			if viper.GetString("prometheus") == "" {
				er(`prometheus needs to be specified. 
				Use --prometheus flag or create config under path $HOME/.promag.yaml`)
			}
			url := promhandler.BuildRulesUrl()
			json := parse.GetJson(url)
			parse.ParseJson(json)
			_ = viper.WriteConfig()
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.promadg.yaml)")
	rootCmd.PersistentFlags().StringP("template", "t", "", "template for alert doc generation")
	rootCmd.PersistentFlags().StringP("prometheus", "p", "", "prometheus url path")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.BindPFlag("prometheus", rootCmd.PersistentFlags().Lookup("prometheus"))
	viper.BindPFlag("template", rootCmd.PersistentFlags().Lookup("template"))

	rootCmd.AddCommand(versionCmd)
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			er(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
