package main

import (
	"fmt"
	"sort"
)

// "math/rand"

type User struct {
	Email string
	Name  string
}

func getEmail(u *User) string {
	return u.Email
}

func getEmail2(u User) string {
	return u.Email
}

// changing the state, needs the pointer.
func updateEmail(u *User, email string) {
	u.Email = email
}

func updateEmail2(u User, email string) {
	u.Email = email
}

func main() {
	names := [3]string{"amir", "mahkam", "ala"}
	secondName := &names[1]

	fmt.Println(*secondName)
	sort.Strings(names[:])
	fmt.Println(*secondName)
}
