// http://127.0.0.1:2379/pd/api/v1/stores

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	url    string
	pdurls = "pd/api/v1/stores"
)

func init() {
	getTiDBInfoCMD.PersistentFlags().StringVarP(&url, "pd leader url", "", "", "pd leader url")
}

var getStoreCMD = &cobra.Command{
	Use:   "getStore",
	Short: "ti-cli get tidb cluster store info",
	Long:  `ti-cli get all tidb cluster store info by rest interface`,
	RunE:  getStore,
}

func getStore(_ *cobra.Command, args []string) error {
	showStoreHost()
	return nil
}

func showStoreHost() error {

	body, status, err := httpGetPd(pdurls)
	if err != nil {
		return err
	}
	if status != http.StatusOK {
		fmt.Println("status=", status)
		fmt.Println(string(body))
		return nil
	}

	var data map[string]interface{}
	json.Unmarshal(body, &data)

	fmt.Println("this tidb cluster store count is : ", data["count"])
	storeinfo := data["stores"].([]interface{})
	fmt.Println("store id and status address relation and tikv status is :")
	for _, item := range storeinfo {
		storeinfo := item.(map[string]interface{})
		st := storeinfo["store"].(map[string]interface{})
		fmt.Printf("%v, %v, %v \n", st["id"], st["status_address"], st["state_name"])
	}

	return nil
}

func httpGetPd(path string) (body []byte, status int, err error) {
	url := "http://" + pdHost.String() + ":" + strconv.Itoa(int(pdPort)) + "/" + path
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
