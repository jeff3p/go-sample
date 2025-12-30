
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

    // HTML content with Tailwind CSS for styling.
    // The content is a single, centered <h1> tag with very large, bold text.
    htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Project Hummingbird</title>
    <!-- Load Tailwind CSS from CDN for instant styling -->
    <script src="https://cdn.tailwindcss.com"></script>
    <link href="https://fonts.googleapis.com/css2?family=Inter:wght@400;700;800&display=swap" rel="stylesheet">
    <style>
        body {
            font-family: 'Inter', sans-serif;
        }
    </style>
</head>
<body class="bg-gray-50 flex items-center justify-center min-h-screen p-4">
    <div class="text-center">
        <h1 class="text-6xl md:text-8xl font-extrabold text-indigo-700 
                   hover:scale-105 transition duration-300 ease-in-out">
            Welcome to Project Hummingbird
        </h1>
        <p class="mt-4 text-xl text-gray-500">
            Your simple Go server is running!
        </p>
        <p class="mt-5 text-xl text-gray-400">
            ....everybody loves hummingbirds
        </p>
    </div>
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
