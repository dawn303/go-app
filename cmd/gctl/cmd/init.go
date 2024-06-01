package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func initConfig() {
	if cfgFile != "" {
		fmt.Println("Use config file: ", cfgFile)
		viper.SetConfigFile(cfgFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Please specify config file")
	rootCmd.PersistentFlags().StringP("name", "n", "Go-app", "Please specify name")
	viper.BindPFlag("name", rootCmd.PersistentFlags().Lookup("name"))
}
