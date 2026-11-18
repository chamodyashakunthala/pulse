package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/chamodyashakunthala/pulse/internal/event"
)

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MB limit
	var e event.Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if err := e.Validate(time.Now().UTC()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO (next step): publish e to RabbitMQ.
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/events", eventsHandler)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	log.Println("ingest listening on :8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
