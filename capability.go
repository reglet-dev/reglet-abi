package abi

import "time"

// ContextWire is the JSON wire format for context.Context propagation.
type ContextWire struct {
	Deadline  *time.Time `json:"deadline,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
	TimeoutMs int64      `json:"timeout_ms,omitempty"`
	Canceled  bool       `json:"canceled,omitempty"`
}
