package routes

import (
	"net/http"
	"text/template"
)

// Templates - Global Variable of templates
var Templates *template.Template

// InitTemplates - init	Templates
func InitTemplates() error {
	files, err := template.ParseFiles("internal/web/public/index.html", "internal/web/public/error.html")
	if err != nil {
		return err
	}
	Templates = template.Must(files, nil)
	return nil
}

// RenderTemplate - Render template by template name
func RenderTemplate(w http.ResponseWriter, tmpl string, vars interface{}) {
	err := Templates.ExecuteTemplate(w, tmpl+".html", vars)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
