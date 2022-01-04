package helper

import (
	"log"
	"net/http"
)

// LogHandle - logging handler
func LogHandle(handlerName string, r *http.Request) {
	log.Printf("%s\t%v\t%v\t%v", handlerName, r.Method, r.Host, r.URL.Path)
}
