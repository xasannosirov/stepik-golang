package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	// "io"
	// "net/url"
	// "os"
	// "time"
)

var counter int

func Handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, counter)
	case http.MethodPost:

		err := r.ParseForm()
		if err != nil {
			fmt.Println("error")
		}

		v := r.Form.Get("count")
		num, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("это не число"))
			return
		}

		counter += int(num)
	default:
		http.Error(w, "method is not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/count", Handler)

	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal()
	}
}
