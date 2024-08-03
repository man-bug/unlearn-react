package main

import (
    "log"
    "net/http"
    "os"

    "github.com/man-bug/unlearn-react/internal/handlers"
)

func main() {
    log.Println("Application starting...")

    // Serve static files
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    // Set up routes
    http.HandleFunc("/", handlers.HomeHandler)
    http.HandleFunc("/increment", handlers.IncrementHandler)

    // Get port from environment variable
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port if not specified
        log.Println("No PORT environment variable detected, defaulting to 8080")
    }

    // Log the port we're about to listen on
    log.Printf("Attempting to start server on port %s", port)

    // Start the server
    log.Printf("Server starting on :%s", port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
