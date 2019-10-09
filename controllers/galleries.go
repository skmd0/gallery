package controllers

import (
	"gallery/models"
	"gallery/views"
)

func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		New: views.NewView("bootstrap", "galleries/new"),
		gs:  gs,
	}
}

// Galleries is a wrapper over View struct used for rendering the template
type Galleries struct {
	New *views.View
	gs  models.GalleryService
}
