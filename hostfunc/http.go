package hostfunc

import "time"

// ContextWire is the JSON wire format for context.Context propagation.
type ContextWire struct {
	Deadline  *time.Time `json:"deadline,omitempty"`
	RequestID string     `json:"request_id,omitempty"`
	TimeoutMs int64      `json:"timeout_ms,omitempty"`
	Canceled  bool       `json:"canceled,omitempty"`
}

// HTTPRequest is the JSON wire format for an HTTP request.
type HTTPRequest struct {
	Headers map[string][]string `json:"headers,omitempty"`
	Method  string              `json:"method"`
	URL     string              `json:"url"`
	Body    string              `json:"body,omitempty"`
	Context ContextWire         `json:"context"`
}

// HTTPResponse is the JSON wire format for an HTTP response.
type HTTPResponse struct {
	Headers       map[string][]string `json:"headers,omitempty"`
	Error         *ErrorDetail        `json:"error,omitempty"`
	Body          string              `json:"body,omitempty"`
	Proto         string              `json:"proto,omitempty"`
	StatusCode    int                 `json:"status_code"`
	BodyTruncated bool                `json:"body_truncated,omitempty"`
}
