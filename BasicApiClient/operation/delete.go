package operation

import (
	"api-client/config"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Delete() {
	url := config.AccountUrl()
	var deleteUrl = url + "/eb0bd6f5-c3f5-44b2-b677-acd23cdde514?version=0"

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", deleteUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
	fmt.Println("response body : ", respBody)
}
