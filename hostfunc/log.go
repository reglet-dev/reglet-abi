package hostfunc

import "time"

// LogMessage is the JSON wire format for a log message from Guest to Host.
type LogMessage struct {
	Timestamp time.Time   `json:"timestamp"`
	Level     string      `json:"level"`
	Message   string      `json:"message"`
	Attrs     []LogAttr   `json:"attrs,omitempty"`
	Context   ContextWire `json:"context"`
}

// LogAttr represents a single slog attribute for wire transfer.
type LogAttr struct {
	Key   string `json:"key"`
	Type  string `json:"type"`  // "string", "int64", "bool", "float64", "time", "error", "any"
	Value string `json:"value"` // String representation of the value
}
