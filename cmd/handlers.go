package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	datatable := scrape("https://www.saihduero.es/risr/EM171")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datatable)
}

func reservoir(w http.ResponseWriter, r *http.Request) {

	datatable := scrape("https://www.saihduero.es/risr/EM171")

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datatable[id])
}
