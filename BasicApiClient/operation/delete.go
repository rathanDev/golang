package operation

import (
	"api-client/config"
	"fmt"
	"net/http"
)

func Delete(id string, version int64) {

	client := &http.Client{}

	url := config.AccountUrl()
	const (
		urlTemplate = `%s/%s?version=%d`
	)
	var deleteUrl = fmt.Sprintf(urlTemplate, url, id, version)
	fmt.Println(deleteUrl)

	// var deleteUrl := createDeleteUrl(id, version)

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

	// printResponse(resp)
	fmt.Println("response Status : ", resp.Status)
	fmt.Println("response Headers : ", resp.Header)
}

func createDeleteUrl(id string, version int64) string {
	url := config.AccountUrl()
	const (
		urlTemplate = `%s/%s?version=%s`
	)
	var deleteUrl = fmt.Sprintf(urlTemplate, url, id, version)
	fmt.Println(deleteUrl)
	return deleteUrl
}

// func printResponse(resp *http.Response) {
// 	fmt.Println("response Status : ", resp.Status)
// 	fmt.Println("response Headers : ", resp.Header)
// }
