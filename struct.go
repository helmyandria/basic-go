package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var data Customer
	data.Name = "Go"
	data.Address = "IND"
	data.Age = 17

	fmt.Println(data)
	fmt.Println(data.Name)
	fmt.Println(data.Address)
	fmt.Println(data.Age)

	data1 := Customer{
		Name:    "Lang",
		Address: "GER",
		Age:     20,
	}
	fmt.Println(data1)

	data2 := Customer{"GO", "FRN", 21}
	fmt.Println(data2)
}
