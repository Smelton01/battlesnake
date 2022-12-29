package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/smelton01/battlesnake/internal/game"
)

func status() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("STATUS: ", r.URL)
		fmt.Fprintf(w, "live")

	}
}

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

func HandleStart() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := &game.GameState{}

		err := json.NewDecoder(r.Body).Decode(state)
		if err != nil {
			log.Printf("ERROR: Failed to decode start json, %s", err)
			return
		}

		game.Start(state)
	}
}

func HandleMove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := &game.GameState{}

		err := json.NewDecoder(r.Body).Decode(&state)
		if err != nil {
			log.Printf("ERROR: Failed to decode move json, %s", err)
			return
		}

		response := game.Move(state)

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("ERROR: Failed to encode move response, %s", err)
			return
		}
	}
}

func HandleEnd() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := &game.GameState{}

		err := json.NewDecoder(r.Body).Decode(state)
		if err != nil {
			log.Printf("ERROR: Failed to decode end json, %s", err)
			return
		}
		game.End(state)
	}
}
