package routes

import (
	"log"
	"net/http"
)

// Should Use? Or use Async Js?
type IndexPage struct {
	Input        string
	Result       string
	ErrorMessage string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("IndexHandler\t%v\t%v", r.Method, r.URL.Path)
	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
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
