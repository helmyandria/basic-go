package main

import "fmt"

func main() {

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	texts := []string{"Java", "Go", "Rust"}
	for i := 0; i < len(texts); i++ {
		fmt.Println(texts[i])
	}

	for index, name := range texts {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range texts {
		fmt.Println(name)

	}

}
