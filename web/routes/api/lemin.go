package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	lemin "lem-in/lemin"
)

func LeminHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("API LeminHandler\t%v\t%v", r.Method, r.URL.Path)
	if r.URL.Path != "/api/lemin" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET")
	case http.MethodPost:
		// Take query

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
		// Start Match result
		// result, err := lemin.GetResultByContent(data.Content, 100_000)
		w.WriteHeader(http.StatusOK)
		err = lemin.WriteResultByContent(data.Content, w)
		// fmt.Printf("%+v\n###\n", w)
		if err != nil {
			log.Print(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%v", err.Error())
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
