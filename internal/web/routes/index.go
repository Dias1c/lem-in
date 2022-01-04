package routes

import (
	"net/http"

	helper "github.com/Dias1c/lem-in/internal/web/helper"
)

// IndexHandler - Main Page handler
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// log.Printf("IndexHandler\t%v\t%v", r.Method, r.URL.Path)
	helper.LogHandle("IndexHandler", r)
	if r.URL.Path != "/" {
		RenderErrorPage(w, "", http.StatusText(http.StatusNotFound))
		return
	}
	// Render
	switch r.Method {
	case http.MethodGet:
		// Here is Default
		RenderTemplate(w, "index", nil)
	case http.MethodPost:
		// Here is Logic
		RenderTemplate(w, "index", nil)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
