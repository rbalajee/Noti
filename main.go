package main

import (
	"expvar"
	"fmt"
	"log"
	"net/http"

	"github.com/rbalajee/noti/handlers"
)

const port = ":3001"

func main() {

	// Create a New ServeMux for routing requests
	mux := http.NewServeMux()

	// Convert our APIListener function to a Handler using mux.HandleFunc which implements the ServeHTTP method
	mux.HandleFunc("/listen", handlers.APIListener)

	mux.HandleFunc("/debug/vars", expvar.Handler().ServeHTTP)

	// Start a server and listen to connections on port 3001
	fmt.Printf("\nListening to connections on path `/listen` on port %s\n", port)
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("\nERROR: Unable to start the listener :: ERROR MESSAGE: %v", err)
	}

}
