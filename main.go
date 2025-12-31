
package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
    })

    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
