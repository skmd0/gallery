package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	layoutDir   = "views/layouts/"
	templateExt = ".gohtml"
)

// NewView builds a View struct
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// View is used to render the correct gohtml files
type View struct {
	Template *template.Template
	Layout   string
}

// Render is used to render the view with the predefined layout
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func layoutFiles() []string {
	files, err := filepath.Glob(layoutDir + "*" + templateExt)
	if err != nil {
		panic(err)
	}
	return files
}