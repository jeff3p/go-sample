package main

import (
	"log"
	"net/http"
 "time"
 "os"
)

const port = ":8080"

func handler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	log.Printf("REQ  %s %s from %s ua=%q", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())

	// Respond based on the exact path
	switch r.URL.Path {
	case "/":
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, err := w.Write([]byte("document root"))
		if err != nil {
			log.Printf("Error writing response for /: %v", err)
		}

	case "/banana":
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, err := w.Write([]byte("banana"))
		if err != nil {
			log.Printf("Error writing response for /banana: %v", err)
		}

	case "/watermelon":
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, err := w.Write([]byte("watermelon"))
		if err != nil {
			log.Printf("Error writing response for /banana: %v", err)
		}

	default:
		// Custom 404: set status and write message
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("404 not found"))
		if err != nil {
			log.Printf("Error writing 404 response: %v", err)
		}
	}

	log.Printf("Handled in %v", time.Since(start))
}


func main() {
	http.HandleFunc("/", handler)

	log.Printf("Starting server on port %s\n", port)
	log.Println("Access it at http://localhost" + port)

  sample := os.Getenv("SAMPLE")
  log.Println("Environment variable loaded" + sample)

  sample2 := os.Getenv("SAMPLE2")
  log.Println("Environment variable loaded" + sample2)

  sample3 := os.Getenv("SAMPLE3")
  log.Println("Environment variable loaded" + sample3)

	log.SetFlags(log.LstdFlags)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
