package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Go")
	fmt.Println(result)

	fmt.Println(helper.Application) //berhasil diakses di luar class
	// fmt.Println(helper.version) //gagal diakses di luar class
}
