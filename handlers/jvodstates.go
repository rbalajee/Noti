package handlers

import (
	"encoding/json"
	_ "expvar"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

// Fields that will be read from the request body when unmarshalling to output
type JVOD struct {
	JobID     string `json:"job_id"`
	Event     string `json:"event"`
	JvodState string `json:"jvod_state"`
}

func JVODStateChangeListener(w http.ResponseWriter, r *http.Request) {

	// Read the request body using io.ReadAll
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("\n ERROR: Error reading the request body %v", err)
	}

	var out JVOD

	// JSON unmarshall the request body to UnmarshalledOutput struct
	_ = json.Unmarshal(body, &out)

	/*
		if err != nil {
			log.Printf("\nERROR: Unmarshalling data from request body :: ERROR MESSAGE: %v", err)
		}
	*/

	if out.Event == "" {
		return
	} else {
		fmt.Println("\n=== JVOD State Change Registered ===")
		fmt.Printf("\nEvent: %s", out.Event)
		fmt.Printf("\nJVOD ID: %s", out.JobID)
		fmt.Printf("\nJVOD State: %s", out.JvodState)
		fmt.Printf("\nTime: %v\n", time.Now())
	}
}
