package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rbalajee/noti/routes"
)

func main() {

	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("\nERROR: Loading the .env file :: ERROR MESSAGE: %v", err)
	}

	// Get port variable from .env file
	port, ok := os.LookupEnv("port")
	if !ok {
		log.Fatalln("ERROR: port variable is not available in .env")
	}

	// Initialize port in case it is set to a zero value string
	if port == "" {
		port = ":3001"
	}

	// Start a net/http server and listen to connections
	fmt.Printf("\nListening to connections on path `/listen` on port `%s`\n", port)
	err = http.ListenAndServe(port, routes.Routes())
	if err != nil {
		log.Fatalf("\nERROR: Unable to start the listener :: ERROR MESSAGE: %v", err)
	}

}
