package handlers

import (
    "html/template"
    "net/http"
)

var (
    // Parse templates. Once, not on every request. Efficiency, heard of it?
    templates = template.Must(template.ParseFiles(
        "internal/templates/layout.html",
        "internal/templates/home.html",
    ))
    counter = 0 // Our mighty counter. React would've used useState for this.
)

// HomeHandler handles the home page. No useEffect needed.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "layout.html", counter)
}

// IncrementHandler handles incrementing the counter. HTMX magic incoming.
func IncrementHandler(w http.ResponseWriter, r *http.Request) {
    counter++ // Increment like a boss.
    w.Header().Set("Content-Type", "text/html")
    templates.ExecuteTemplate(w, "counter", counter) // Partial template update!
}
