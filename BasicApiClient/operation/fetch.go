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

func MyAccount() model.Account {
	var account model.Account
	account.Id = "001"
	account.Name = "Jana"
	account.Country = "Singapore"
	return account
}

func Fetch() model.Account { //model.ApiResponse.Accounts {

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
	fmt.Printf("apiResponse.AccountDataList: %s", apiResponse.AccountDataList)

	fmt.Println()

	accountDataList := apiResponse.AccountDataList
	fmt.Println("\n")
	fmt.Println("accountDataList => ", accountDataList)

	printAccounts(accountDataList)

	var account model.Account
	account.Id = "001"
	account.Name = "Jana"
	account.Country = "Singapore"
	return account

	return account
}

func printAccounts(accountDataList []model.AccountData) {
	for i, val := range accountDataList {
		fmt.Println("i =>", i)

		fmt.Println("Attributes =>", *val.Attributes)
		fmt.Println("ID =>", val.ID)
		fmt.Println("OrganisationID =>", val.OrganisationID)
		fmt.Println("Type =>", val.Type)
		fmt.Println("Version =>", *val.Version)
	}
}
