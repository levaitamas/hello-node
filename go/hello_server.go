package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const defPort int = 8080

type FixCharReadSeeker struct {
	numBytes int64
	char     byte
}

func NewFixCharReadSeeker(numBytes int64, char byte) FixCharReadSeeker {
	return FixCharReadSeeker{numBytes, char}
}

func (f FixCharReadSeeker) Read(p []byte) (n int, err error) {
	for i := 0; i < len(p); i++ {
		p[i] = f.char
	}
	return len(p), nil
}

func (f FixCharReadSeeker) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		return offset, nil
	case io.SeekCurrent:
		return 0, nil
	case io.SeekEnd:
		return f.numBytes - offset, nil
	default:
		return -1, fmt.Errorf("invalid 'whence' parameter")
	}
}

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

		http.ServeContent(w, r, "", time.Time{},
			NewFixCharReadSeeker(int64(numBytes), byte('x')))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request for URL:", r.URL)
		fmt.Fprintf(w, "Hello World from %s!\n", hostname)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("Error (http server): ", err)
	}

}
