package cmd

import (
	"fmt"
	"net"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	host   net.IP
	port   uint16
	genDoc bool
	pdHost net.IP
	pdPort uint16
	// ca        string
	// sslCert   string
	// sslKey    string
	ctlClient *http.Client
	// schema    string
	user   string
	passwd string
)

var rootCmd = &cobra.Command{
	Use:   "ti-cli",
	Short: "ti-cli controller",
	Long:  `ti-cli is a command tool for tidb dba`,
	RunE:  genDocument,
}

func genDocument(c *cobra.Command, args []string) error {
	if !genDoc || len(args) != 0 {
		return c.Usage()
	}
	docDir := "./doc"
	docCmd := &cobra.Command{
		Use:   "ti-cli",
		Short: "ti-cli controller",
		Long:  `ti-cli is a command tool for tidb dba`,
	}
	docCmd.AddCommand(getTiDBInfoCMD)
	fmt.Println("Generating documents...")
	if err := doc.GenMarkdownTree(docCmd, docDir); err != nil {
		return err
	}
	fmt.Println("Done!")
	return nil
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.AddCommand(getTiDBInfoCMD)

	rootCmd.PersistentFlags().IPVarP(&host, "host", "", net.ParseIP("172.16.7.150"), "TiDB server port")
	rootCmd.PersistentFlags().Uint16VarP(&port, "port", "", 10080, "TiDB server status port")
	rootCmd.PersistentFlags().IPVarP(&pdHost, "pdHost", "", net.ParseIP("172.16.7.150"), "PD server port")
	rootCmd.PersistentFlags().Uint16VarP(&pdPort, "pdPort", "", 2379, "PD server status port")
	rootCmd.PersistentFlags().StringVarP(&user, "user", "", "", "user name")
	rootCmd.PersistentFlags().StringVarP(&passwd, "password", "", "", "user passwd")
	rootCmd.Flags().BoolVar(&genDoc, "gendoc", false, "generate doc file")
	if err := rootCmd.Flags().MarkHidden("gendoc"); err != nil {
		fmt.Printf("can not mark hidden flag, flag %s is not found", "gendoc")
		return
	}
	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(initCmd)

	cobra.OnInitialize(func() {
		ctlClient = &http.Client{
			Transport: &http.Transport{},
		}
	})

}
