package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var activeConnections int32

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the GoLang API server!")
	time.Sleep(10 * time.Second)
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running!")
}

func connectionsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Active Connections: %d", atomic.LoadInt32(&activeConnections))
}

func connectionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&activeConnections, 1)
		defer atomic.AddInt32(&activeConnections, -1)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/status", statusHandler)
	mux.HandleFunc("/connections", connectionsHandler)

	loggedMux := connectionMiddleware(mux)

	fmt.Println("Starting server on port 8080...")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatal(err)
	}
}
