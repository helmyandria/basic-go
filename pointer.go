package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Surabaya", "Jawa Timur", "Indonesia"}
	address2 := address1 //copy value

	address2.City = "Jombang"
	fmt.Println(address1)
	fmt.Println(address2)

	// pointer : create reference ke lokasi data di memory yg sama, tanpa duplikasi data yg sudah ada
	addressPointer1 := Address{"Semarang", "Jawa Tengah", "Indonesia"}
	addressPointer2 := &addressPointer1 //copy value

	addressPointer2.City = "Klaten"
	fmt.Println(addressPointer1)
	fmt.Println(addressPointer2)
}
