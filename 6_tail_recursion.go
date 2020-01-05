package main

import "fmt"

// Note: using tail recursion the same function can be written as below,
// but Go doesn't optimize this, though there are workarounds, still it performed better in benchmarks.
// For better performance, we are using tail recursion.
func factorialTailRec(num int) int {
	return factorial(1, num)
}

func factorial(accumulator, val int) int {
	if val == 1 {
		return accumulator
	}
	return factorial(accumulator*val, val-1)
}

func main() {

	fmt.Println(factorialTailRec(5)) // 120
}
