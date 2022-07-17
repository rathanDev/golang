package main

import (
	"api-client/operation"
	"fmt"
)

func main() {
    accounts := operation.Fetch()
	fmt.Println(accounts)
}
