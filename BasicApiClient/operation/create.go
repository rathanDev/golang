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

	var jsonPayload = []byte(`{
		"data": {
		  "type": "accounts",
		  "id": "ad27e265-9605-4b4b-a0e5-3003ea9cc504",
		  "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde604",
		  "attributes": {
			"country": "GB",
			"base_currency": "GBP",
			"bank_id": "400300",
			"bank_id_code": "GBDSC",
			"bic": "NWBKGB22",
			"name": [
			  "Janarthan"
			],
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
	`)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("response:", resp)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
}

func createNumberPointer(x int64) *int64 {
	return &x
}
