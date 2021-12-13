// Always start source files with package declaration

package models

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
