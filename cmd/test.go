package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "version",
	Short: "ti-cli version 0.1",
	Long:  `this is ti-cli 0.1 , it's a test cli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ti-cli 0.1")
	},
}
