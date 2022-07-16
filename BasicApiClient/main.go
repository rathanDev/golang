package main

import (
	"api-client/helper"
	"api-client/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {

	url := helper.GetUrl()

	response, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	fmt.Println()

	accountData := model.AccountData{ID: "someId", OrganisationID: "orgId"}
	fmt.Println(accountData)

	fmt.Println("")

    jsonstr := `[{"GBP":657.54},{"USD":123.45}]`
    fmt.Println(jsonstr)

    ps := []model.Price{}
    json.Unmarshal([]byte(jsonstr), &ps)
    fmt.Println(ps)

}
