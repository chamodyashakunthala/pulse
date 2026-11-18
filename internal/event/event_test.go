package event

import (
	"testing"
	"time"
)

func TestValidate(t *testing.T) {
	now := time.Date(2026, 11, 16, 8, 0, 0, 0, time.UTC)

	tests := []struct {
		name    string
		event   Event
		wantErr bool
	}{
		{"valid event", Event{Service: "payments", LatencyMs: 85, Status: 200}, false},
		{"missing service", Event{LatencyMs: 85, Status: 200}, true},
		{"negative latency", Event{Service: "payments", LatencyMs: -5, Status: 200}, true},
		{"absurd latency", Event{Service: "payments", LatencyMs: 700000, Status: 200}, true},
		{"bad status", Event{Service: "payments", LatencyMs: 85, Status: 42}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.event.Validate(now)
			if (err != nil) != tc.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}

func TestValidateFillsTimestamp(t *testing.T) {
	now := time.Date(2026, 11, 16, 8, 0, 0, 0, time.UTC)
	e := Event{Service: "payments", LatencyMs: 85, Status: 200}
	if err := e.Validate(now); err != nil {
		t.Fatal(err)
	}
	if !e.Timestamp.Equal(now) {
		t.Errorf("expected timestamp %v, got %v", now, e.Timestamp)
	}
}
