package hostfunc

// SMTPRequest is the JSON wire format for an SMTP connection request.
type SMTPRequest struct {
	Host      string      `json:"host"`
	Port      string      `json:"port"`
	Context   ContextWire `json:"context"`
	TimeoutMs int         `json:"timeout_ms,omitempty"`
	TLS       bool        `json:"tls"`
	StartTLS  bool        `json:"starttls"`
}

// SMTPResponse is the JSON wire format for an SMTP connection response.
type SMTPResponse struct {
	Error          *ErrorDetail `json:"error,omitempty"`
	Address        string       `json:"address,omitempty"`
	Banner         string       `json:"banner,omitempty"`
	TLSVersion     string       `json:"tls_version,omitempty"`
	TLSCipherSuite string       `json:"tls_cipher_suite,omitempty"`
	TLSServerName  string       `json:"tls_server_name,omitempty"`
	ResponseTimeMs int64        `json:"response_time_ms,omitempty"`
	Connected      bool         `json:"connected"`
	TLS            bool         `json:"tls,omitempty"`
}
