package operation

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

func Greet() string {
	var greeting = "Fetch operation says Hi"
	return greeting
}

func Fetch() {

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
	fmt.Println("responseData => ", string(responseData))

	fmt.Println()

	var apiResponse model.ApiResponse
	json.Unmarshal([]byte(string(responseData)), &apiResponse)
	fmt.Println("apiRespone => ", apiResponse, "\n")
	fmt.Printf("apiResponse.Data: %s", apiResponse.Accounts)

	fmt.Println()

	// var accounts []AccountData
	accounts := apiResponse.Accounts
	fmt.Println("\n")
	fmt.Println("accounts => ", accounts)


	// for id, organisation_id := range accounts {
	for i, val := range apiResponse.Accounts {
		// Each value is an interface{} type, that is type asserted as a string
		// fmt.Println(id, organisation_id)
		fmt.Println("i =>", i)
		fmt.Println("value =>", val)

		fmt.Println("ID =>", val.ID)
		fmt.Println("OrganisationID =>", val.OrganisationID)
		fmt.Println("Type =>", val.Type)
		fmt.Println("Version =>", val.Version)
	}

	// return accounts
}
