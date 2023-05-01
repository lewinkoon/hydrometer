package main

import (
	"encoding/json"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	datatable := Scrape("https://www.saihduero.es/risr/EM171")

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datatable)
}
