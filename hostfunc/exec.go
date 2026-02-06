package hostfunc

// ExecRequest is the JSON wire format for an exec request.
type ExecRequest struct {
	Args    []string    `json:"args"`
	Env     []string    `json:"env,omitempty"`
	Command string      `json:"command"`
	Dir     string      `json:"dir,omitempty"`
	Context ContextWire `json:"context"`
}

// ExecResponse is the JSON wire format for an exec response.
type ExecResponse struct {
	Error      *ErrorDetail `json:"error,omitempty"`
	Stdout     string       `json:"stdout"`
	Stderr     string       `json:"stderr"`
	ExitCode   int          `json:"exit_code"`
	DurationMs int64        `json:"duration_ms,omitempty"`
	IsTimeout  bool         `json:"is_timeout,omitempty"`
}
