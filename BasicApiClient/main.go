package main

import (
	"api-client/model"
	"api-client/operation"
	"api-client/util"
	"fmt"
)

func main() {

	testCreate()

	// operation.Create()

	// operation.Delete()
}

func testCreate() {
	var accountData model.AccountData
	accountData.ID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde516"
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = util.CreateNumberPointer(1)

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = util.CreateStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = util.CreateBooleanPointer(false)
	accountAttr.AccountNumber = "1006"
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC"
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = util.CreateStringPointer("SG")
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
