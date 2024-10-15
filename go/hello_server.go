package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Error (query hostname): ", err)
		os.Exit(1)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request for URL:", r.URL)
		fmt.Fprintf(w, "Hello World from %s!", hostname)
	})
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error (http server): ", err)
	}

}
