package main

import (
	"api-client/model"
	"api-client/operation"
	"fmt"
)

func main() {

	testCreate()

	// operation.Create()

	// operation.Delete()
}

func testCreate() {
	var accountData model.AccountData
	accountData.ID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde515"
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde615"
	accountData.Type = "accounts"
	accountData.Version = operation.CreateNumberPointer(1)

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = operation.CreateStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = operation.CreateBooleanPointer(false)
	accountAttr.AccountNumber = "1005"
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC"
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = operation.CreateStringPointer("SG")
	accountAttr.Name = []string{"Jana", "Param"}

	accountData.Attributes = &accountAttr

	operation.Create(accountData)
}

func tryFetch() {
	fmt.Println("Fetch")
	apiResponse := operation.Fetch()
	fmt.Println(apiResponse)

	fmt.Println("\n\nFetchMapped")
	accounts := operation.FetchMapped()
	fmt.Println(accounts)
}
