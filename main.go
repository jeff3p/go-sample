
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    // Disable timestamps in log output
    log.SetFlags(0)

    // Startup logs (no time)
    log.Println("Starting HTTP server on :8080")
    log.Println("Access at http://localhost:8080")

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Per-request log (before handling)
        log.Printf("REQ %s %s from %s ua=%q", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())

        w.Header().Set("Content-Type", "text/plain; charset=utf-8")

        // Read env vars inline
        sample := os.Getenv("SAMPLE")
        if sample == "" {
            sample = "(unset)"
        }
        sample2 := os.Getenv("SAMPLE2")
        if sample2 == "" {
            sample2 = "(unset)"
        }

        switch r.URL.Path {
        case "/":
            fmt.Fprintf(w, "document root\nSAMPLE=%s\nSAMPLE2=%s\n", sample, sample2)
        case "/banana":
            fmt.Fprintf(w, "banana\nSAMPLE=%s\nSAMPLE2=%s\n", sample, sample2)
        default:
            http.NotFound(w, r)
        }

        // Completion log (no duration)
        log.Printf("DONE %s %s", r.Method, r.URL.Path)
    })

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Server failed: %v", err)
    }
}
