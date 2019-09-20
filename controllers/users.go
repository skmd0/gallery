package controllers

import (
	"gallery/views"
	"net/http"
)

// NewUsers is used to create anew Users controller
// this function will panic if the templates are not
// parsed correclty and should only be used during the
// initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
