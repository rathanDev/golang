package main

import (
	"api-client/operation"
	"fmt"
)

func main() {
	fmt.Println("Fetch")
	accounts := operation.Fetch()
	fmt.Println(accounts)

	fmt.Println("\n\nFetchWhole")
	//var apiResponse model.ApiResponse
	apiResponse := operation.FetchWhole()
	fmt.Println(apiResponse)
}
