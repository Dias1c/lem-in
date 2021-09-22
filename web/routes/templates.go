package routes

import (
	"net/http"
	"text/template"
)

var Templates *template.Template

func InitTemplates() error {
	files, err := template.ParseFiles("web/public/index.html")
	if err != nil {
		return err
	}
	Templates = template.Must(files, nil)
	return nil
}

func RenderTemplate(w http.ResponseWriter, tmpl string, vars interface{}) {
	err := Templates.ExecuteTemplate(w, tmpl+".html", vars)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}
