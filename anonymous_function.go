package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "anjing"
	}

	registerUser("Go", blacklist) //method call 1

	registerUser("anjing", func(name string) bool {
		return name == "anjing"
	}) //method call 2

}
