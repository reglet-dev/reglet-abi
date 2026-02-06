package hostfunc

// DNSRequest is the JSON wire format for a DNS lookup request.
type DNSRequest struct {
	Hostname   string      `json:"hostname"`
	Type       string      `json:"type"`
	Nameserver string      `json:"nameserver,omitempty"`
	Context    ContextWire `json:"context"`
}

// DNSResponse is the JSON wire format for a DNS lookup response.
type DNSResponse struct {
	Error     *ErrorDetail `json:"error,omitempty"`
	Records   []string     `json:"records,omitempty"`
	MXRecords []MXRecord   `json:"mx_records,omitempty"`
}

// MXRecord represents a single MX record.
type MXRecord struct {
	Host string `json:"host"`
	Pref uint16 `json:"pref"`
}
