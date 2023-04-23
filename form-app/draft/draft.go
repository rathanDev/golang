package draft

import (
	"encoding/json"
	"log"
	"net/http"

	"form3-client/config"
	"form3-client/model"
	"form3-client/operation"
	"form3-client/util"
)

func non_main() {

	log.Println("#--- Form3-client - Start as service ---#")

	http.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("\nExecute create fetch and delete")
		writer.Header().Set("Content-Type", "application/json")
		jsonResponse := TestCreateFetchDelete()
		writer.Write(jsonResponse)
		return
	})

	log.Fatal(http.ListenAndServe(":8085", nil))
}

func TestCreateFetchDelete() []byte {

	resp := make(map[string]string)

	const localhostUrl = "http://localhost:8080"
	config.SetBaseUrl(localhostUrl)

	const givenAccountNumber = "400300"
	const givenAccountId = "eb0bd6f5-c3f5-44b2-b677-acd23cdde516"
	var givenVersion = util.CreateNumberPointer(0)

	var accountData model.AccountData
	accountData.ID = givenAccountId
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = givenVersion

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = util.CreateStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = util.CreateBooleanPointer(false)
	accountAttr.AccountNumber = givenAccountNumber
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC"
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = util.CreateStringPointer("SG")
	accountAttr.Name = []string{"Jana", "Param"}

	accountData.Attributes = &accountAttr

	operation.Create(accountData)

	accounts := operation.FetchMapped()
	log.Println("AccountsAfterCreation:", accounts)

	if len(accounts) == 1 {
		resp["ExpectOneAccountToBeCreated"] = "PASS"
	} else {
		resp["ExpectOneAccountToBeCreated"] = "FAIL"
	}

	account := accounts[0]
	var fetchedAccountNumber = account.AccountNumber

	if fetchedAccountNumber == givenAccountNumber {
		resp["ExpectFetchAccountNumberIsEquealToGivenAccountNumber"] = "PASS"
	} else {
		resp["ExpectFetchAccountNumberIsEquealToGivenAccountNumber"] = "FAIL"
	}

	operation.Delete(givenAccountId, *givenVersion)

	accountsAfterDeletion := operation.FetchMapped()
	log.Println("AccountsAfterDeletion:", accountsAfterDeletion)

	if len(accountsAfterDeletion) == 0 {
		resp["ExpectNoAccounts"] = "PASS"
	} else {
		resp["ExpectNoAccounts"] = "FAIL"
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	return jsonResp
}

func create() {
	log.Println("\n\nCreate")

	var accountData model.AccountData
	accountData.ID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde516"
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde616"
	accountData.Type = "accounts"
	accountData.Version = util.CreateNumberPointer(0)

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

func fetch() {
	log.Println("\n\nFetch")
	apiResponse := operation.Fetch()
	log.Println(apiResponse)

	log.Println("\nFetchMapped")
	accounts := operation.FetchMapped()
	log.Println(accounts)
}

func delete() {
	log.Println("\n\nDelete")
	var id = "eb0bd6f5-c3f5-44b2-b677-acd23cdde516"
	var version = int64(0)
	operation.Delete(id, version)
}
