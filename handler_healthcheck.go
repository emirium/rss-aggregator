package main

import "net/http"

func handlerHealthcheck(w http.ResponseWriter, r *http.Request) {
	// return 200 OK
	respondWithJSON(w, 200, struct{}{})
}
