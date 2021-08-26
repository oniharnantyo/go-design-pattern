package main

import (
	"./person"
	"fmt"
)

func main()  {
	pb := person.NewPersonBuilder()
	pb.
		Lives().At("Jalan solo").In("Yogyakarta").WithPostCode("23123").
		Works().At("Company").AsA("Programmer").Earning(10000000)
	person := pb.Build()
	fmt.Println(person)
}

