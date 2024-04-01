package main

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover()                  // catch data panic, and continue process from panic
	fmt.Println("Terjadi panic", message) // pemanfaatan recover harus di dalam defer
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
}

func main() {
	runApp(true)
	fmt.Println("running...")
}
