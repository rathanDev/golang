package main

import (
	"fmt"
	"structs-example/model"
)


type Person struct {
	name string
	age  int
}

func main() {
	p := Person{"Jana", 31}
	fmt.Println(p)

	// type Account a

	a := model.Account{"601", "DBS"}
	fmt.Println(a)
}
