package hostfunc

import "time"

// TCPRequest is the JSON wire format for a TCP connection request.
type TCPRequest struct {
	Host      string      `json:"host"`
	Port      string      `json:"port"`
	Context   ContextWire `json:"context"`
	TimeoutMs int         `json:"timeout_ms,omitempty"`
	TLS       bool        `json:"tls"`
}

// TCPResponse is the JSON wire format for a TCP connection response.
type TCPResponse struct {
	TLSCertNotAfter *time.Time   `json:"tls_cert_not_after,omitempty"`
	Error           *ErrorDetail `json:"error,omitempty"`
	TLSVersion      string       `json:"tls_version,omitempty"`
	LocalAddr       string       `json:"local_addr,omitempty"`
	TLSCipherSuite  string       `json:"tls_cipher_suite,omitempty"`
	TLSServerName   string       `json:"tls_server_name,omitempty"`
	TLSCertSubject  string       `json:"tls_cert_subject,omitempty"`
	TLSCertIssuer   string       `json:"tls_cert_issuer,omitempty"`
	RemoteAddr      string       `json:"remote_addr,omitempty"`
	Address         string       `json:"address,omitempty"`
	ResponseTimeMs  int64        `json:"response_time_ms,omitempty"`
	TLS             bool         `json:"tls,omitempty"`
	Connected       bool         `json:"connected"`
}
