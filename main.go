package main

import (
	"encoding/json"
	"net/http"
)

// Response represents a response sent back to user
type Response struct {
	IP      string
	Headers http.Header
}

func printHeader(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := Response{r.RemoteAddr, r.Header}
	json.NewEncoder(w).Encode(response)

}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	if true {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
}

func main() {
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/", printHeader)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
