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

func Create() {
	url := config.AccountUrl()

	var accountData model.AccountData
	accountData.ID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde512"
	accountData.OrganisationID = "eb0bd6f5-c3f5-44b2-b677-acd23cdde612"
	accountData.Type = "accounts"

	var accountAttr model.AccountAttributes
	accountAttr.AccountClassification = createStringPointer("Personal")
	accountAttr.AccountMatchingOptOut = createBooleanPointer(false)
	accountAttr.AccountNumber = "1001"
	accountAttr.AlternativeNames = []string{"Jana", "Rathan"}
	accountAttr.BankID = "400300"
	accountAttr.BankIDCode = "GBDSC"
	accountAttr.BaseCurrency = "SGD"
	accountAttr.Bic = "NWBKGB22"
	accountAttr.Country = createStringPointer("SG")
	accountAttr.Name = []string{"Jana", "Param"}

	accountData.Attributes = &accountAttr

	payload := createPayloadV(accountData)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("response:", resp)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
}

func createPayloadV(accountData model.AccountData) []byte {
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

func createBooleanPointer(x bool) *bool {
	return &x
}

func createStringPointer(x string) *string {
	return &x
}

func createNumberPointer(x int64) *int64 {
	return &x
}
