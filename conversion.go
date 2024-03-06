package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16) // overflow limit int16

	var text = "Belajar Go Lang"
	var g uint8 = text[8]
	var gString = string(g)

	fmt.Println(text)
	fmt.Println(g)
	fmt.Println(gString)
}
