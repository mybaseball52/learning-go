package controllers

import (
	"net/http"
)

func RegisterControllers() {
	uc := newUserCtrl()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
