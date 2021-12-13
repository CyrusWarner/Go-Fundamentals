package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController() // gets the instance of the userController

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

}
