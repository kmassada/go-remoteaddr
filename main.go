package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Ip      string
	Headers http.Header
}

func printHeader(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	response := Response{r.RemoteAddr, r.Header}
	json.NewEncoder(w).Encode(response)

}
func main() {
	http.HandleFunc("/", printHeader)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
