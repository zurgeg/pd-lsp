package lsd

type Message struct {
	text string
}

type RequestMessage struct {
	Message
	id     interface{}
	method string
	params interface{}
}

type ResponseMessage struct {
	Message
	id     interface{}
	result interface{}
	error_ ResponseError
}

type ResponseError struct {
	code    int
	message string
	data    interface{}
}

const (
	ParseError                 = -32700
	InvalidRequest             = -32600
	MethodNotFound             = -32601
	InvalidParams              = -32602
	InternalError              = -32603
	jsonrpcReservedErrorsStart = -32099
	// Deprecated: Use jsonrpcReservedErrorsStart
	serverErrorStart = jsonrpcReservedErrorsStart
)
