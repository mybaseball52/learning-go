package controllers

import (
	"net/http"
	"regexp"
)

type userCtrl struct {
	userIDPattern *regexp.Regexp
}

func (uc userCtrl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from User Ctrl"))
}

func newUserCtrl() *userCtrl {
	return &userCtrl{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
