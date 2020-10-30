package resultok

// ResultOk is a struct, which will returns as a good request result
type ResultOk struct {
	Status string
}

// NewResult returns a new object of Result struct
func NewResult(status string) *ResultOk {
	return &ResultOk{Status: status}
}
