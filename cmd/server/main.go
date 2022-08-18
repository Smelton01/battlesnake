package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/smelton01/battlesnake/internal/api"
)

func main() {
	r := chi.NewRouter()

	api.Route(r)
	r.Get("/status", status())

	log.Printf("Listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func status() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "live")
	}
}
