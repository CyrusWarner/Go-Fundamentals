// Always start source files with package declaration

package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

// variable block
var (
	users  []*User // slice that holds pointers to User objects
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id")
	}
	u.ID = nextID
	nextID++                  // increments the nextID
	users = append(users, &u) // appends the memory location address to the users pointer array
	return u, nil
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id) // fmt.Errorf allows us to pass id into the '%v' to display the not found user id
}

func UpdateUser(u User) (User, error) {
	for i, user := range users {
		if user.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserByID(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...) // Takes all the indexes before and not including index i, then we add the indexes after index i but not including index i removing that user from the users
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)

}
