package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Promadg",
	Long:  `All software has versions. This is Promadg's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Promadg v0.1 -- HEAD")
	},
}
