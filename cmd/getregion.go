package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	titabs    = "tables"
	tiregions = "regions"
)

func init() {
	//getRegionCMD.PersistentFlags().StringVarP(&url, "tihost", "ph", "", "tidb status url")
}

var getRegionCMD = &cobra.Command{
	Use:   "getRegions",
	Short: "ti-cli get tidb cluster region info",
	Long:  `ti-cli get all tidb cluster region info by table name`,
	RunE:  getRegion,
}

//url := "http://" + host.String() + ":" + strconv.Itoa(int(port)) + "/" + path
// curl http://{TiDBIP}:10080/tables/{db}/{table}/regions
// curl http://{TiDBIP}:10080/regions/{regionID}
func getRegion(_ *cobra.Command, args []string) error {
	url := "http://" + host.String() + ":" + strconv.Itoa(int(port)) + "/" + titabs + "/" + db + "/" + table + "/" + tiregions
	body, status, err := httpGetRegion(url)
	if err != nil {
		return err
	}
	if status != http.StatusOK {
		// Print response body directly if status is not ok.
		fmt.Println("host or port maybe wrong")
		return nil
	}
	var data map[string]interface{}
	json.Unmarshal(body, &data)

	fmt.Printf("table name is : %v，table id is : %v \n", data["name"], data["id"])
	regions := data["record_regions"].([]interface{})
	for _, item := range regions {
		region := item.(map[string]interface{})
		region_id := region["region_id"]
		leader := region["leader"].(map[string]interface{})
		fmt.Printf("region id is : %v, store id is : %v \n", region_id, leader["store_id"])
	}

	indices := data["indices"].([]interface{})
	if len(indices) == 0 {
		fmt.Println("have no index")
	} else {
		for _, idx := range indices {
			idxmap := idx.(map[string]interface{})
			idxname := idxmap["name"]
			idxid := idxmap["id"]
			fmt.Printf("index name is : %v, index id is : %v \n", idxname, idxid)

			iregions := idxmap["regions"].([]interface{})

			for _, idxregion := range iregions {
				iregion := idxregion.(map[string]interface{})
				iregion_id := iregion["region_id"]
				getRegionInfo(iregion_id)
				istoreid := iregion["leader"].(map[string]interface{})
				fmt.Printf("index region id is : %v, store id is: %v \n", iregion_id, istoreid["store_id"])
			}

		}
	}

	return nil
}

func getRegionInfo(region interface{}) (err error) {
	region_id := strconv.FormatFloat(region.(float64), 'f', -1, 64)
	url := "http://" + host.String() + ":" + strconv.Itoa(int(port)) + "/" + tiregions + "/" + region_id
	body, status, err := httpGetRegion(url)
	if err != nil {
		return err
	}
	if status != http.StatusOK {
		// Print response body directly if status is not ok.
		fmt.Println("host or port maybe wrong")
		return nil
	}
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	fmt.Printf("start key is : %v,end key is: %v \n", data["start_key"], data["end_key"])

	frames := data["frames"].([]interface{})

	for _, item := range frames {
		itemjson := item.(map[string]interface{})
		if itemjson["is_record"] == false {
			fmt.Printf("index region ,it db name is : %v,table name is: %v,table id is: %v ,index name is: %v ,index id is: %v \n", itemjson["db_name"], itemjson["table_name"], itemjson["table_id"], itemjson["index_name"], itemjson["index_id"])

		} else {
			fmt.Printf("table region ,db name is : %v,table name is: %v,table id is: %v \n", itemjson["db_name"], itemjson["table_name"], itemjson["table_id"])
		}
	}

	return nil
}
