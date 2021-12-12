package routes

import "net/http"

// ErrorPage - using for Render Error Page
type ErrorPage struct {
	Title   string
	Message string
}

// RenderErrorPage - Rendering error page for client
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
