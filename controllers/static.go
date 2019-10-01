package controllers

import (
	"gallery/views"
)

// NewStatic returns a struct will static view objects
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

// Static contains simple static view objects
type Static struct {
	Home    *views.View
	Contact *views.View
}
