package main

import "fmt"

// If we pass parameters by the value we can ensure referential transparency even
// if there is an accidental mutation of passed data within the function
func main() {
	type Person struct {
		firstName string
		lastName  string
		fullName  string
		age       int
	}
	var getFullName = func(in Person) string {
		in.fullName = in.firstName + in.lastName
		return in.fullName
	}

	john := Person{
		"john", "doe", "", 30,
	}

	fmt.Println(getFullName(john)) // johndoe
	fmt.Println(john)              // {john doe  30}
}
