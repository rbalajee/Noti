package routes

import (
	"expvar"
	"net/http"

	"github.com/rbalajee/noti/handlers"
)

func Routes() *http.ServeMux {

	// Create a New ServeMux for routing requests
	mux := http.NewServeMux()

	// Convert our APIListener function to a Handler using mux.HandleFunc which implements the ServeHTTP method
	mux.HandleFunc("/listen", handlers.APIListener)

	mux.HandleFunc("/debug/vars", expvar.Handler().ServeHTTP)

	return mux

}
