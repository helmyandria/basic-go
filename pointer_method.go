package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() { //send pointer in method param
	man.Name = "Mr. " + man.Name
}

func main() {

	name := Man{"Go"}
	name.Married()

	fmt.Println(name.Name)
}
