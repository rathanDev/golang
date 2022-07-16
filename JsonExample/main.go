package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string
	Description string
}

type AccountData struct {
	ID             string `json:"id,omitempty"`
	OrganisationID string `json:"organisation_id,omitempty"`
}

type ApiResponse struct {
	Data []AccountData `json:"data,omitempty"`
}

func main() {

	// accountJson := `{"data":[{"id":"f50fc252-03f3-11ed-a0bc-0242ac130011","modified_on":"2022-02-03T00:00:00.000Z","organisation_id":"f50fc252-03f3-11ed-a0bc-0242ac130012","type":"accounts","version":1},{"attributes":{"account_classification":"class1","alternative_names":null,"country":"Singapore","name":["Jana"]},"created_on":"2022-01-03T00:00:00.000Z","id":"f50fc252-03f3-11ed-a0bc-0242ac130013","modified_on":"2022-02-03T00:00:00.000Z","organisation_id":"f50fc252-03f3-11ed-a0bc-0242ac130014","type":"accounts","version":1},{"attributes":{"account_classification":"class1","account_matching_opt_out":true,"account_number":"1O001","alternative_names":null,"bank_id":"5O001","country":"Singapore","name":["Jana"]},"created_on":"2022-01-03T00:00:00.000Z","id":"f50fc252-03f3-11ed-a0bc-0242ac130101","modified_on":"2022-02-03T00:00:00.000Z","organisation_id":"f50fc252-03f3-11ed-a0bc-0242ac130201","type":"accounts","version":1}],"links":{"first":"/v1/organisation/accounts?page%5Bnumber%5D=first","last":"/v1/organisation/accounts?page%5Bnumber%5D=last","self":"/v1/organisation/accounts"}}`
	accountJson := `{"data":[ {"id":"id1","organisation_id":"orgId1","type":"accounts"}, {"id":"id2","organisation_id":"orgId2"} ]}`
	var apiResponse ApiResponse
	json.Unmarshal([]byte(accountJson), &apiResponse)
	fmt.Println("apiRespone => ", apiResponse)
	fmt.Printf("apiResponse.Data: %s", apiResponse.Data)

	fmt.Println()

	// var accounts []AccountData
	accounts := apiResponse.Data
	fmt.Println("accounts => ", &accounts)

	for id, organisation_id := range accounts {
		// Each value is an interface{} type, that is type asserted as a string
		fmt.Println(id, organisation_id)
	}
}

func unmarshallBird() {
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)
	fmt.Println()
	fmt.Println("bird => ", bird)

	birdsJson := `[{"species": "pigeon","description": "likes to perch on rocks"}, {"species": "mina","description": "likes to live with people"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdsJson), &birds)
	fmt.Println("birds => ", birds)
}
