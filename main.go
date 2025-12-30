
package main

import (
    "fmt"
    "log"
    "net/http"
)

const port = ":8080"

func handler(w http.ResponseWriter, r *http.Request) {
    // Log each request
    log.Printf("Request: method=%s path=%s remote=%s ua=%q",
        r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())

    // Only serve the root path
    if r.URL.Path != "/" {
        // Log the 404
        log.Printf("Not Found: path=%s", r.URL.Path)

        // Optional: custom 404 HTML
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        w.WriteHeader(http.StatusNotFound)
        _, _ = w.Write([]byte(`
<!DOCTYPE html>
<html lang="en">
  <head><meta charset="utf-8"><title>404 Not Found</title></head>
  <body>
    <h1>404 Not Found</h1>
    <p>The requested path was not found.</p>
  </body>
</html>
`))
        return
    }

    // Normal response for "/"
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    w.WriteHeader(http.StatusOK)

    htmlContent := `
<!DOCTYPE html>
<html lang="en">
  <head><meta charset="utf-8"><title>Go Server</title></head>
  <body>
    <p class="mt-4 text-xl text-gray-500">
      Your simple Go server is running!
    </p>
  </body>
</html>
`
    if _, err := w.Write([]byte(htmlContent)); err != nil {
        log.Printf("Error writing response: %v", err)
    }
}

func main() {
    http.HandleFunc("/", handler)

    fmt.Printf("Starting server on port %s\n", port)
    fmt.Println("Access it at http://localhost" + port)

    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
