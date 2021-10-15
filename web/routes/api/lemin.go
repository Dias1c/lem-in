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
		fmt.Fprintf(w,
			`This path accepts only the POST method. 
The data that this page accepts must be in json format.

JSON = "
{
	Content: "Data"
}
"

Data = "
AntsCount (Integer > 0)
Rooms (roomName coorX coorY) + Start and End rooms
Relations (roomName-roomName)
"
`)
	case http.MethodPost:
		// Take query

		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		var data = struct{ Content string }{}
		err := dec.Decode(&data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		// Start Match result
		if lErr := lemin.WriteResultByContent(w, data.Content); lErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%v", lErr.Error())
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
}
