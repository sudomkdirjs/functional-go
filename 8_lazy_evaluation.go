package main

import "fmt"

// Note: We can use higher-order-functions to rewrite this into a lazily evaluated version

func add(x int) int {
	fmt.Println("executing add")
	return x + x
}

func multiply(x int) int {
	fmt.Println("executing multiply")
	return x * x
}

func main() {
	fmt.Println(addOrMultiply(true, add, multiply, 4))
	fmt.Println(addOrMultiply(false, add, multiply, 4))
}

// This is now a higher-order-function hence evaluation of the functions are delayed in if-else
func addOrMultiply(add bool, onAdd, onMultiply func(t int) int, t int) int {
	if add {
		return onAdd(t)
	}
	return onMultiply(t)
}

// Output
// executing add
// 8
// executing multiply
// 16

// Note: There are also other ways of doing it using Sync & Futures like this and using channels and goroutines like this.
// Doing Lazy evaluations in Go might not be worth the code complexity most of the times, but if the functions in question are
// heavy in terms of processing then its is absolutely worth it to lazy evaluate them.
