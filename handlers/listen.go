package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Fields that will be read from the request body when unmarshalling to output
type UnmarshalledOutput struct {
	Event string `json:"event"`
	Job   struct {
		ID    string `json:"id"`
		State string `json:"state"`
	} `json:"job"`
}

func APIListener(w http.ResponseWriter, r *http.Request) {

	// Read the request body using io.ReadAll
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("\n ERROR: Error reading the request body %v", err)
	}

	var out UnmarshalledOutput

	// JSON unmarshall the request body to UnmarshalledOutput struct
	err = json.Unmarshal(body, &out)
	if err != nil {
		log.Printf("\nERROR: Unmarshalling data from request body :: ERROR MESSAGE: %v", err)
	}

	fmt.Printf("\nEvent: %s", out.Event)
	fmt.Printf("\nJob ID: %s", out.Job.ID)
	fmt.Printf("\nJob State: %s\n", out.Job.State)
}
