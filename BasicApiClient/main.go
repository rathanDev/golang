package main

import (
	"api-client/operation"
	"fmt"
)

func main() {
	fmt.Println("Fetch")
	apiResponse := operation.Fetch()
	fmt.Println(apiResponse)

	fmt.Println("\n\nFetchMapped")
	accounts := operation.FetchMapped()
	fmt.Println(accounts)
}
