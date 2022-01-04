package web

import (
	"log"
	"net/http"
	"time"

	routes "github.com/Dias1c/lem-in/internal/web/routes"
	routesApi "github.com/Dias1c/lem-in/internal/web/routes/api"
)

// RunServer - starts server with setted port
func RunServer(port string) error {
	var err error
	err = ValidatePort(port)
	if err != nil {
		return err
	}
	err = routes.InitTemplates()
	if err != nil {
		return err
	}

	// Init Handlers + Run Server
	mux := http.NewServeMux()
	// FS
	fs := http.FileServer(http.Dir("internal/web/public/assets/"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
	maps := http.FileServer(http.Dir("maps/"))
	mux.Handle("/maps/", http.StripPrefix("/maps/", maps))
	// Pages
	mux.HandleFunc("/", routes.IndexHandler) // Index Page
	// APIs
	mux.HandleFunc("/api/lemin", routesApi.LeminHandler)
	// Start Listen
	server := http.Server{
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	log.Printf("Server started on http://localhost%v", port)
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
