package operation

import (
	"api-client/config"
	"api-client/util"
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

	util.PrintHttpResponse(resp)
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
