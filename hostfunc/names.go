// Package hostfunc defines the wire types for Reglet host functions.
package hostfunc

const (
	ModuleName    = "reglet_host"
	HTTPRequestOp = "http_request"
	DNSLookupOp   = "dns_lookup"
	TCPConnectOp  = "tcp_connect"
	SMTPConnectOp = "smtp_connect"
	ExecCommandOp = "exec_command"
)
