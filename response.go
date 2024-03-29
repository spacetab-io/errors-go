package errors

type (
	FieldName       string
	ValidationError string
	ErrorCode       string
	ErrorType       string
	DebugData       interface{}
)

type Response struct {
	Error ErrorObject `json:"error"`
}

const (
	ErrorTypeError   ErrorType = "error"
	ErrorTypeWarning ErrorType = "warning"
	ErrorTypeInfo    ErrorType = "info"
)

type ErrorObject struct {
	Message    interface{}                     `json:"message"`
	Type       *ErrorType                      `json:"type,omitempty"`
	Code       *ErrorCode                      `json:"code,omitempty"`
	Validation map[FieldName][]ValidationError `json:"validation,omitempty"`
	Debug      *DebugData                      `json:"debug,omitempty"`
}
