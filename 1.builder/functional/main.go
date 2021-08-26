package main

import (
	"fmt"

	"./person"
)

func main() {
	b := person.PersonBuilder{}
	p := b.Called("Person").WorksAsA("Programmer").Build()
	fmt.Println(p)
}
