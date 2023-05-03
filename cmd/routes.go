package main

import (
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()
	mux.Get("/", app.home)
	mux.Get("/level/", app.level)

	return mux

}
