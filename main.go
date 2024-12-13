package main

import (
    "fmt"
    "net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/health" {
        w.WriteHeader(http.StatusOK)
        fmt.Fprintln(w, `{"status": "healthy"}`)
    } else {
        w.WriteHeader(http.StatusNotFound)
        fmt.Fprintln(w, `{"error": "not found"}`)
    }
}

func main() {
    http.HandleFunc("/", healthCheckHandler)
    fmt.Println("Server is starting on port 8080...")
    http.ListenAndServe(":8080", nil)
}
