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

func Fetch() []model.Account {

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

	var apiResponse model.ApiResponse
	json.Unmarshal([]byte(string(responseData)), &apiResponse)

	accountDataList := apiResponse.AccountDataList

	var accounts []model.Account

	for i, accountData := range accountDataList {
		fmt.Println("i =>", i)

		var account model.Account
		account.ID = accountData.ID
		account.OrganisationID = accountData.OrganisationID
		account.Type = accountData.Type

		account.AccountNumber = accountData.Attributes.AccountNumber
		account.BankID = accountData.Attributes.BankID
		account.BankIDCode = accountData.Attributes.BankIDCode
		account.Country = *accountData.Attributes.Country
		account.Name = accountData.Attributes.Name

		accounts = append(accounts, account)
	}

	return accounts
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
