package main

import (
	"fmt"
)

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
