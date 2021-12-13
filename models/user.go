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

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++                  // increments the nextID
	users = append(users, &u) // appends the memory location address to the users pointer array
	return u, nil
}
