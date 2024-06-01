package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version short desc",
	Long:  "version long desc",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("gctl version: v1.0.0")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
