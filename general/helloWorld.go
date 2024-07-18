package main

import "fmt"

var helloVar string = "Hello, world!"

const helloConst string = "Hello, world!"

var (
	name string = "Misha"
	age  int    = 24
)

func main() {
	fmt.Println(helloVar)
	fmt.Println(helloConst)
	fmt.Printf("My name is %s, and I'm %d years old.\n", name, age)
}
