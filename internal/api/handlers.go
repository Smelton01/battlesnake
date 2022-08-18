package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/smelton01/battlesnake/internal/game"
)

func HandleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := game.Info()

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("ERROR: Failed to encode response, %s", err)
			return
		}
	}
}
