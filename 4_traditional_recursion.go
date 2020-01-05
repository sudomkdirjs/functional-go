package main

import "fmt"

func factorial(num int) int {

	result := 1

	for ; num > 0; num-- {
		result *= num
	}

	return result
}

func main() {

	fmt.Println(factorial(5)) // 120
}
