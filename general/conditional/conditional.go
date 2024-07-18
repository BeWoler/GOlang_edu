package main

import "fmt"

func main() {
	const a int = 5
	const b int = 10

	if a >= 5 {
		fmt.Println(a)
	} else {
		fmt.Println("a must be greater than 5")
	}

	switch b {
	case 5:
		fmt.Println("b is 5")
	case 10:
		fmt.Println("b is 10")
	default:
		fmt.Println("b is neither 5 nor 10")
	}
}
