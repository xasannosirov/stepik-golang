package main

import (
	"fmt"
	"io"
	"net/http"
	// "os"
	// "time"
	"log"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:5555/get")

	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf(string(data))
}
