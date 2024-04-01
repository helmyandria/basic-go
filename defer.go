package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	defer logging() // function dijadwalkan di eksekusi setelah sebuah function di eksekusi
	fmt.Println("Run application")
}

func main() {
	runApplication()
}
