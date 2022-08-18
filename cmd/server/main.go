package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/smelton01/battlesnake/internal/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	r := chi.NewRouter()

	api.Route(r)
	r.Get("/status", status())

	log.Printf("Listening on port %s.", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func status() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "live")

	}
}
