package event

import (
	"errors"
	"strings"
	"time"
)

// Event is one request that a monitored service reports to Pulse.
type Event struct {
	Service   string    `json:"service"`
	Version   string    `json:"version"`
	LatencyMs int       `json:"latency_ms"`
	Status    int       `json:"status"`
	Timestamp time.Time `json:"ts"`
}

const maxLatencyMs = 600_000 // 10 minutes; anything larger is almost surely a bug

// Validate checks the event is sane and fills in Timestamp if it is missing.
func (e *Event) Validate(now time.Time) error {
	e.Service = strings.TrimSpace(e.Service)
	if e.Service == "" {
		return errors.New("service is required")
	}
	if len(e.Service) > 64 {
		return errors.New("service name too long (max 64)")
	}
	if e.LatencyMs < 0 || e.LatencyMs > maxLatencyMs {
		return errors.New("latency_ms must be between 0 and 600000")
	}
	if e.Status < 100 || e.Status > 599 {
		return errors.New("status must be a valid HTTP status (100-599)")
	}
	if e.Timestamp.IsZero() {
		e.Timestamp = now
	}
	return nil
}
