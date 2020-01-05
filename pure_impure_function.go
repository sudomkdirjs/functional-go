package main

import "fmt"

// pure function will always return the same output for the given input
// and its behavior is highly predictable.
func pureSum(a, b int) int {
	return a + b
}

var holder = map[string]int{}

// If we add an extra line in the pure function,
// the behavior becomes unpredictable as it now has a side effect that affects an external state.
func impureSum(a, b int) int {
	c := a + b
	holder[fmt.Sprintf("%d+%d", a, b)] = c
	return c
}

func main() {

	fmt.Println(pureSum(5, 10))

	fmt.Println(holder)

	fmt.Println(impureSum(5, 10))

	fmt.Println(holder)

}
