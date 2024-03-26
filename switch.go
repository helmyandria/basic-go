package main

import "fmt"

func main() {
	text := "Go"

	switch text {
	case "Go":
		fmt.Println("Hello Go")
	case "Java":
		fmt.Println("Hello Java")
	default:
		fmt.Println("Default Value")
	}

	switch length := len(text); length > 5 { //switch short statement
	case true:
		fmt.Println("incorrect")
	case false:
		fmt.Println("correct")
	}
}
