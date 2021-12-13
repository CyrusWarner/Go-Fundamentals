package controllers // package declaration

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHttp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}

func newUserController() *userController { // Convention typically used for creating a constructor function
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(/d+)/?`),
	}
}