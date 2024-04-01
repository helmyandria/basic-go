package main

import "fmt"

func endApp() {
	fmt.Println("End app")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR") //stop program dan tetap eksekusi defer
	}
}

func main() {
	runApp(true)
}
