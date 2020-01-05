package main

import "fmt"

// Note: The downside of the recursive approach is that it will be slower
// compared to an iterative approach most of the times(The advantage we are aiming for is code simplicity
// and readability) and might result in stack overflow errors since every function call
// needs to be saved as a frame to the stack. To avoid this tail recursion is preferred, especially when
// the recursion is done too many times. In tail recursion, the recursive call is the last thing executed by the function and
// hence the functions stack frame need not be saved by the compiler.
// Most compilers can optimize the tail recursion code the same way iterative code is optimized hence avoiding the performance penalty.
// Go compiler, unfortunately, does not do this optimization.

func factorial(num int) int {

	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func main() {

	fmt.Println(factorial(5)) // 120
}
