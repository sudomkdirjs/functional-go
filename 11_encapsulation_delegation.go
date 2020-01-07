package main

import (
	"fmt"
)

// Adding methods to a non-struct type (like an int or string) might be a code smell that you are unnecessarily leaking implementation details or other logic that should be private.
// If the details need to be hidden, methods can still be added to the custom type, while at the same time hiding where that logic comes from.

// Delegation is used inside the User struct. The details of how the ID calculates old school is even hidden to User.

// The implementation details (that old school access is based on the ID) can be hidden from the rest of the code.

// internally, the User struct will take advantage of the methods on the ID, further encapsulating the logic.

const LastOldSchoolID = 23959

type UserID int

func (id UserID) IsOldSchool() bool {
	return id < LastOldSchoolID // The last oldschool player
}

func (id UserID) IsNotOldSchool() bool {
	return !id.IsOldSchool()
}

type User struct {
	UserID
}

func (u User) IsOldSchool() bool {
	return u.UserID.IsOldSchool()
}

func (u User) IsNotOldSchool() bool {
	return u.UserID.IsNotOldSchool()
}

func main() {
	u := User{1}
	fmt.Printf("User can access the oldschool game: %v \n", u.IsOldSchool())
	fmt.Printf("User cannot access the oldschool game: %v \n", u.IsNotOldSchool())

	u = User{100000}
	fmt.Printf("User can access the oldschool game: %v \n", u.IsOldSchool())
	fmt.Printf("User cannot access the oldschool game: %v \n", u.IsNotOldSchool())
}
