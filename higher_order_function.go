package main

import "fmt"

func main() {
	var list = []string{"Orange", "Apple", "Banana", "Grape"}

	var getItemLengthFunc = func(item string) int {
		return len(item)
	}

	// we are passing the array and a function as arguments to mapForEach method.
	var out = mapForEach(list, getItemLengthFunc)

	fmt.Println(out)
}

// The higher-order-function takes an array and a function as arguments
func mapForEach(arr []string, fn func(item string) int) []int {

	var lenghtArr = []int{}

	for i, item := range arr {
		fmt.Println(i, "=>", item)
		// We are executing the method passed
		lenghtArr = append(lenghtArr, fn(item))
	}

	return lenghtArr
}
