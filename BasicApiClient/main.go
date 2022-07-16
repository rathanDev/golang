package main

import (
	"api-client/helper"
	"api-client/model"
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

	fmt.Println("")

	accountData := model.AccountData{ID: "someId", OrganisationID: "orgId"}
	fmt.Println(accountData)

}
