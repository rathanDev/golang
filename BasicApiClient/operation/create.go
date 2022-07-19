package operation

import (
	"api-client/config"
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func Create() {
	url := config.AccountUrl()

	var accountType = "accounts"
	var id = "eb0bd6f5-c3f5-44b2-b677-acd23cdde511"
	var orgId = "eb0bd6f5-c3f5-44b2-b677-acd23cdde611"
	var country = "SG"
	var baseCurrency = "SGD"
	var bankId = "400300"
	var bankIdCode = "GBDSC"
	var bic = "NWBKGB22"
	// var name []string = []string{"Jana", "Param"}

	var jsonStr = createPayload(accountType, id, orgId, country, baseCurrency, bankId, bankIdCode, bic)
	var jsonPayload = []byte(jsonStr)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("response:", resp)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
}

func createPayload(accountType string, id string, orgId string, country string, baseCurrency string, bankId string, bankIdCode string, bic string) string {
	const (
		jsonTemplate = `{
			"data": {
			  "type": "%s",
			  "id": "%s",
			  "organisation_id": "%s",
			  "attributes": {
				"country": "%s",
				"base_currency": "%s",
				"bank_id": "%s",
				"bank_id_code": "%s",
				"bic": "%s",
				"name": ["Janarthan Paraman"],
				"alternative_names": [
				  "Jana"
				],
				"account_classification": "Personal",
				"joint_account": false,
				"account_matching_opt_out": false,
				"secondary_identification": "A1B2C3D4"
			  }
			}
		  }
		`
	)
	var jsonStr = fmt.Sprintf(jsonTemplate, accountType, id, orgId, country, baseCurrency, bankId, bankIdCode, bic)
	fmt.Println(jsonStr)
	return jsonStr
}

func createNumberPointer(x int64) *int64 {
	return &x
}
