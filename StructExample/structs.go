package main

import "fmt"

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 31
	return &p
}

func main() {

	p := Person{"Jana", 31}

	fmt.Println(p)

}
