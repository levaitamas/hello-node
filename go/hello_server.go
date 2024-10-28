package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const defPort int = 8080

func main() {
	envPort, ok := os.LookupEnv("PORT")
	port := defPort
	if ok {
		if tmpPort, err := strconv.ParseInt(envPort, 10, 32); err == nil {
			port = int(tmpPort)
		}
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Error (query hostname): ", err)
		os.Exit(1)
	}

	http.HandleFunc("GET /get-bytes/{num_bytes}", func(w http.ResponseWriter, r *http.Request) {
		var numBytes int
		if _, err := fmt.Sscanf(r.PathValue("num_bytes"), "%d", &numBytes); err != nil {
			w.WriteHeader(406)
			return
		}
		if _, err = w.Write(bytes.Repeat([]byte("x"), numBytes)); err != nil {
			log.Println("Error while serving path:", r.URL)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request for URL:", r.URL)
		fmt.Fprintf(w, "Hello World from %s!\n", hostname)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("Error (http server): ", err)
	}

}
