package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "gctl",
	Short: "gctl short desc",
	Long:  "gctl long desc",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("this is RunE function.")
		fmt.Println("RunE: config -> ", viper.Get("cfgFile"))
		fmt.Println("RunE: name -> ", viper.Get("name"))
		fmt.Println("RunE: http addr -> ", viper.Get("server.http.addr"))
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}
}
