package main

import "fmt"

func main() {

	type strung string // strung is name type data alias from string

	var text1 strung = "Belajar"

	var text2 string = "Sinau"
	var text3 strung = strung(text2)

	fmt.Println(text1)
	fmt.Println(text3)
}
