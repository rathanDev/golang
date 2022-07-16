package main

import (
	"api-client/operation"
	"fmt"
)

func main() {

	fmt.Println(operation.Greet())
	fmt.Println()
	fmt.Println()

	operation.Fetch()

}
