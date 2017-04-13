package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		m := &struct {
			Message string `json:"message"`
		}{
			Message: "pong",
		}

		res, _ := json.Marshal(m)
		fmt.Fprintf(w, string(res))
	})
	fmt.Println("http server started on :8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
