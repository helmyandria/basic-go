package main

import "fmt"

func main() {
	const (
		text1 string = "Belajar"
		text2        = "Go Lang"
	)

	fmt.Println(text1)
	fmt.Println(text2)

	// error const cannot assign value to text1
	// text1 = "Sinau"
	// text2 = "Java"
}
