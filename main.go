package main

import (
	"fmt"
	"log"
	"net/http"
)

// The port where the server will listen for requests.
const port = ":8080"

// handler is the HTTP handler function for the root path ("/").
func handler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to tell the browser we are sending HTML.
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	htmlContent := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
    <p class="mt-4 text-xl text-gray-500">
        Your simple Go server is running!
    </p>
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
