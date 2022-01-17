package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

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

func httpGetRegion(url string) (body []byte, status int, err error) {
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
		fmt.Println("host or port maybe wrong")
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
