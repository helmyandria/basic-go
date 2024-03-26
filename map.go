package main

import "fmt"

func main() {

	// var person map[string]string = map[string]string{} //manual declare
	// person["name"] = ""
	// person["address"] = ""

	person := map[string]string{
		"name":    "helmy",
		"address": "Surabaya",
	}
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Clean Code"
	book["author"] = "Robert C. Martin"
	book["ups"] = "Salah"

	fmt.Println(book)

	delete(book, "ups")

	fmt.Println(book)

}
