package user

import (
	"fmt"
	"errors"
)

type User struct {
	firstName string
	lastName string
	birthdate string
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}

func New(firstName string, lastName string, birthdate string) (*User, error) {
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("firstName, lastName and birthdate are required")
	}
	return &User{
		firstName,
		lastName,
		birthdate,
	}, nil
}
