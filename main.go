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
