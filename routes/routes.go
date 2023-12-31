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
	mux.HandleFunc("/state-change", handlers.StateChangeListener)

	mux.HandleFunc("/jvod-state-change", handlers.JVODStateChangeListener)

	mux.HandleFunc("/debug/vars", expvar.Handler().ServeHTTP)

	return mux

}
