package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")

        switch r.URL.Path {
        case "/":
            fmt.Fprint(w, "document root")
        case "/banana":
            fmt.Fprint(w, "banana")
        default:
            // Keep it simple: no logging, default 404
            http.NotFound(w, r)
        }
    })

    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
