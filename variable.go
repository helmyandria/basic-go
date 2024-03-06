package main

import "fmt"

func main() {
	var desc = "Belajar" // declare variable with var
	fmt.Println(desc)

	desc = "Belajar Go Lang"
	fmt.Println(desc)

	text := "Belajar" // declarea variable with :
	fmt.Println(text)

	text = "Belajar Go Lang"
	fmt.Println(text)

	var (
		desc1 = "Belajar"
		desc2 = "Go Lang"
		desc3 = "Dasar"
	)

	fmt.Println(desc1)
	fmt.Println(desc2)
	fmt.Println(desc3)
}
