package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	datatable := scrape("https://www.saihduero.es/risr/EM171/historico/xATSOFURfFzNx0UR")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datatable)
}

func (app *application) level(w http.ResponseWriter, r *http.Request) {

	datatable := scrape("https://www.saihduero.es/risr/EM171/historico/xATSOFURfFzNx0UR")

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(datatable[id])
}
