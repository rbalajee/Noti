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
type UnmarshalledOutput struct {
	Event string `json:"event"`
	Job   `json:"job"`
}

type Job struct {
	ID        string `json:"id"`
	State     string `json:"state"`
	SsaiState string `json:"ssai_state"`
	SEPState  string `json:"sep_state"`
}

func StateChangeListener(w http.ResponseWriter, r *http.Request) {

	// Read the request body using io.ReadAll
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("\n ERROR: Error reading the request body %v", err)
	}

	var out UnmarshalledOutput

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
		fmt.Println("\n=== Job State Change Registered ===")
		fmt.Printf("\nEvent: %s", out.Event)
		fmt.Printf("\nJob ID: %s", out.ID)
		fmt.Printf("\nJob State: %s", out.State)
		fmt.Printf("\nSEP State: %s", out.SEPState)
		fmt.Printf("\nSSAI State: %s", out.SsaiState)
		fmt.Printf("\nTime: %v\n", time.Now())
	}
}
