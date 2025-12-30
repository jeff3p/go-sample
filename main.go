
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

// The port where the server will listen for requests.
const port = ":8080"

// logDir and logFile define where we will write logs
const logDir = "/applog"
const logFile = "requests.log"

// logRequest appends a single line with timestamp and request context.
func logRequest(r *http.Request) {

    // Open the file in append mode; create if it doesn't exist
    path := filepath.Join(logDir, logFile)
    f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
    if err != nil {
        log.Printf("logRequest: failed to open log file: %v", err)
        return
    }
    defer f.Close()

    // Build a simple log line with RFC3339 timestamp and request info
    tstamp := time.Now().Format(time.RFC3339)
    line := fmt.Sprintf("%s\tmethod=%s\tpath=%s\tremote=%s\tua=%q\n",
        tstamp, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())

    // Write the line
    if _, err := f.WriteString(line); err != nil {
        log.Printf("logRequest: failed to write log line: %v", err)
    }
}

// handler is the HTTP handler function for the root path ("/").
func handler(w http.ResponseWriter, r *http.Request) {
    // Log the request first
    logRequest(r)

    // Set the Content-Type header to tell the browser we are sending HTML.
    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
    <p> Your simple Go server is running! </p>
</body>
</html>
`)

    // Write the HTML content to the response writer.
    _, err := w.Write([]byte(htmlContent))
    if err != nil {
        log.Printf("Error writing response: %v", err)
    }
}

func main() {
    // Register the handler function for the root path.
    http.HandleFunc("/", handler)

    fmt.Printf("Starting server on port %s\n", port)
    fmt.Println("Access it at http://localhost" + port)

    // Start the HTTP server. log.Fatal ensures the server exits if it fails to start.
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
