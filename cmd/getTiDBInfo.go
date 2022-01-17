package cmd

import (
	"github.com/spf13/cobra"
)

var (
	status string
)

func init() {
	getTiDBInfoCMD.PersistentFlags().StringVarP(&status, "status", "", "", "tidb status")
}

var getTiDBInfoCMD = &cobra.Command{
	Use:   "getinfo",
	Short: "ti-cli get tidb cluster status info",
	Long:  `ti-cli get all tidb cluster status info by rest interface`,
	RunE:  getInfo,
}

func getInfo(_ *cobra.Command, args []string) error {
	httpPrint("/status")
	return nil
}
