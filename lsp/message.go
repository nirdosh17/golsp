package lsp

type Request struct {
	RPC    string `json:"jsonrpc"` // value always be 2.0
	ID     int    `json:"id"`
	Method string `json:"method"`

	// Implement later
	// Params
}

type Response struct {
	RPC string `json:"jsonrpc"` // value always be 2.0
	ID  *int   `json:"id,omitempty"`

	// Result
	// Error
}

type Notification struct {
	RPC    string `json:"jsonrpc"` // value always be 2.0
	Method string `json:"method"`
}
