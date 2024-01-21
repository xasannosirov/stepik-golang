package main

import (
	"fmt"
	"net/http"
	// "io"
	// "os"
	// "time"
)

func main() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, web!")
	})

	http.ListenAndServe(":8080", nil)
}
