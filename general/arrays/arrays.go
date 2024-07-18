package main

import "fmt"

var numbArr [3]int = [3]int{13, 222, 31}

func main() {
	numbers := [3]int{135, 1, 36}

	dynamicNumbers := [...]int{1, 2, 3}

	fmt.Println(numbArr)
	fmt.Println(numbers)
	fmt.Println(dynamicNumbers)
}
