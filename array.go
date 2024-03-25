package main

import "fmt"

func main() {
	var desc [3]string // declare array

	desc[0] = "Belajar"
	desc[1] = "Go"
	desc[2] = "Lang"

	fmt.Println(desc[0])
	fmt.Println(desc[1])
	fmt.Println(desc[2])

	var values = [...]int{21, 77, 80, 99, 22} // declare and insert array

	fmt.Println(values)
	fmt.Println(len(values))

}
