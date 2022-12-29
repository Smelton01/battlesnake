package api

import "github.com/go-chi/chi/v5"

func BindAll(r *chi.Mux) {
	r.Use(withServerID)

	r.Get("/", HandleIndex())
	r.Post("/start", HandleStart())
	r.Post("/move", HandleMove())
	r.Post("/end", HandleEnd())
	r.Get("/status", status())
}
