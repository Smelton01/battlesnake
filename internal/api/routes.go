package api

import "github.com/go-chi/chi/v5"

func Route(r *chi.Mux) {
	r.Use(withServerID)
	r.Get("/", HandleIndex())
	r.Get("/start", HandleStart())
	r.Get("/move", HandleMove())
	r.Get("/end", HandleEnd())

}
