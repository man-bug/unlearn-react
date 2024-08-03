package main

import (
    "fmt"
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

    // Start the server, bind to all interfaces
    log.Printf("Server starting on 0.0.0.0:%s", port)
    err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
