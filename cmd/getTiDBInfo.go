package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
	Short: "ti-cli get tidb info",
	Long:  `ti-cli get all tidb info by rest interface`,
	RunE:  getInfo,
}

func getInfo(_ *cobra.Command, args []string) error {
	httpPrint("/status")
	return nil
}

func httpGet(path string) (body []byte, status int, err error) {
	url := "http://" + host.String() + ":" + strconv.Itoa(int(port)) + "/" + path
	resp, err := ctlClient.Get(url)
	if err != nil {
		return
	}
	status = resp.StatusCode
	defer func() {
		if errClose := resp.Body.Close(); errClose != nil && err == nil {
			err = errClose
		}
	}()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

func httpPrint(path string) error {
	body, status, err := httpGet(path)
	if err != nil {
		return err
	}
	if status != http.StatusOK {
		// Print response body directly if status is not ok.
		fmt.Println(string(body))
		return nil
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, body, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(prettyJSON.String())
	return nil
}
