package main

import "fmt"

// Lazy evaluation or non-strict evaluation is the process of delaying evaluation of an expression until it is needed.
// In general, Go does strict/eager evaluation but for operands like && and || it does a lazy evaluation.
// We can utilize higher-order-functions, closures, goroutines, and channels to emulate lazy evaluations.

func main() {
	fmt.Println(addOrMultiply(true, add(4), multiply(4)))  // 8
	fmt.Println(addOrMultiply(false, add(4), multiply(4))) // 16
}

func add(x int) int {
	fmt.Println("executing add") // this is printed since the functions are evaluated first
	return x + x
}

func multiply(x int) int {
	fmt.Println("executing multiply") // this is printed since the functions are evaluated first
	return x * x
}

func addOrMultiply(add bool, onAdd, onMultiply int) int {
	if add {
		return onAdd
	}
	return onMultiply
}

// Output
// executing add
// executing multiply
// 8
// executing add
// executing multiply
// 16
