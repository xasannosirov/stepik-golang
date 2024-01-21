package main

import (
	"fmt"
	"net/http"
    "log"
	//"io"
	// "os"      
	// "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
    
	message := fmt.Sprintf("Hello,%s!", name)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/api/user", handler)
    err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}