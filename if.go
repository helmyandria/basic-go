package main

import "fmt"

func main() {
	text := "Java"

	if text == "Go" {
		fmt.Println("Hello Go")
	} else if text == "Java" {
		fmt.Println("Hello Java")
	} else {
		fmt.Println("Exception statement")
	}

	if length := len(text); length > 5 { //if short statement ;
		fmt.Println("text terlalu panjang")
	} else {
		fmt.Println("text correct")
	}
}
