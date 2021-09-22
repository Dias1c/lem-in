package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	lemin "lem-in/lemin"
)

func LeminHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("LeminHandler\t%v\t%v", r.Method, r.URL.Path)
	if r.URL.Path != "/api/lemin" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET")
	case http.MethodPost:
		// Take query

		log.Printf("Body: %+v\n", r.Body)
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		var data = struct{ Content string }{}
		err := dec.Decode(&data)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		log.Printf("data: %+v\n", data)
		// Start Match result
		result, err := lemin.GetResultByContent(data.Content)
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%v", err.Error())
			// http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// time.Sleep(time.Second * 5)
		// Set the header and send data
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%v", result)
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
