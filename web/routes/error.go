package routes

import "net/http"

type ErrorPage struct {
	Title   string
	Message string
}

func RenderErrorPage(w http.ResponseWriter, title, message string) {
	if title == "" {
		title = "Error"
	}
	if message == "" {
		message = "Something Wrong!"
	}
	errPage := ErrorPage{
		Title:   title,
		Message: message,
	}
	RenderTemplate(w, "error", errPage)
}
