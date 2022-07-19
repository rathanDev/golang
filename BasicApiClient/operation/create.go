package operation

import (
	"api-client/config"
	"api-client/model"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Create(accountData model.AccountData) {
	url := config.AccountUrl()

	payload := createPayload(accountData)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}

	printResponse(resp)
}

func createPayload(accountData model.AccountData) []byte {
	a, _ := json.Marshal(accountData)
	fmt.Println(string(a))
	const (
		jsonTemplate = `{
			"data": %s
		  }
		`
	)
	var jsonPayload = fmt.Sprintf(jsonTemplate, a)
	var payload = []byte(jsonPayload)
	return payload
}

func printResponse(resp *http.Response) {
	fmt.Println("response:", resp)
	fmt.Println("response:", resp)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
}
