package types

// APIError is a type-aliased string that conforms to Error interface
type APIError string

// Conformance to the Error interface
func (a APIError) Error() string {
	return string(a)
}
