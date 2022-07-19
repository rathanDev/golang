package main

import (
	"api-client/operation"
	"fmt"
)

func main() {

	// tryFetch()

	// operation.Create()

	operation.Delete()
}

func tryFetch() {
	fmt.Println("Fetch")
	apiResponse := operation.Fetch()
	fmt.Println(apiResponse)

	fmt.Println("\n\nFetchMapped")
	accounts := operation.FetchMapped()
	fmt.Println(accounts)
}
