package main

import "fmt"

// Note: Unfortunately, there are not many ways to strictly limit data mutation in Go,
// however by using pure functions and by explicitly avoiding data mutations and reassignment using other concepts we saw earlier this can be achieved.
// Go by default passes variables by value, except for slices and maps. So, avoid passing them by reference(using pointers) as much as possible.

// the below will mutate external state as we are passing a parameter by reference and
// hence doesn't ensure referential transparency
func main() {
	type Person struct {
		firstName string
		lastName  string
		fullName  string
		age       int
	}
	var getFullName = func(in *Person) string {
		in.fullName = in.firstName + in.lastName // data mutation
		return in.fullName
	}

	john := Person{
		"john", "doe", "", 30,
	}

	fmt.Println(getFullName(&john)) // johndoe
	fmt.Println(john)               // {john doe johndoe 30}
}
