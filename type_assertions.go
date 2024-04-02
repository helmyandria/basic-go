package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) { // get type interface
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}
