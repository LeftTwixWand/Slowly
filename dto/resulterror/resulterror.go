package resulterror

// ResultError is a struct, which will returns as a bad request result
type ResultError struct {
	Error string `json:"error"`
}

// NewResult returns a new object of ResultError struct
func NewResult(err string) *ResultError {
	return &ResultError{Error: err}
}
