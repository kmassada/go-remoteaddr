package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Response represents a response sent back to user
type Response struct {
	IP      string
	Headers http.Header
}

var started time.Time

func printHeader(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := Response{r.RemoteAddr, r.Header}
	json.NewEncoder(w).Encode(response)

}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	duration := time.Now().Sub(started)
	if duration.Seconds() > 30 {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
}

func main() {
	started = time.Now()
	http.HandleFunc("/healthz", healthzHandler)
	http.HandleFunc("/", printHeader)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
