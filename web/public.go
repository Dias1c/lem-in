package web

import (
	"fmt"
	"log"
	"net/http"
	"time"

	general "github.com/Dias1c/lem-in/general"
	routes "github.com/Dias1c/lem-in/web/routes"
	routesApi "github.com/Dias1c/lem-in/web/routes/api"
)

// RunServer - starts server with setted port
func RunServer(port string) {
	var err error
	err = validatePort(port)
	if err != nil {
		general.CloseProgram(err)
	}
	err = routes.InitTemplates()
	if err != nil {
		general.CloseProgram(err)
	}

	// Init Handlers + Run Server (To Do)
	mux := http.NewServeMux()
	// FS
	assets := http.FileServer(http.Dir("web/public/assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", assets))
	// Pages
	mux.HandleFunc("/", routes.IndexHandler) // Index Page
	// APIs
	mux.HandleFunc("/api/lemin", routesApi.LeminHandler)
	// Start Listen
	addr := fmt.Sprintf("localhost%v", port)
	server := http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Printf("Server started on http://%v", addr)
	err = server.ListenAndServe()
	if err != nil {
		general.CloseProgram(err)
	}
}
