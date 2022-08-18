package api

import "net/http"

const ServerID = "BattlesnakeOfficial/starter-snake-go"

func withServerID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Server", ServerID)
		next.ServeHTTP(w, r)
	})
}
