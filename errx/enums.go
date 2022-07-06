package errx

// Code top-level identifier of an error
type Code int

// SubCode more strict identifier of an error
type SubCode int

var (
	codes    map[string]Code    = make(map[string]Code)
	subCodes map[string]SubCode = make(map[string]SubCode)
)

// RegisterCode registers a new error code
func RegisterCode(namespace string, code int) {
	codes[namespace] = Code(code)
}

// RegisterSubCode registers a new error sub code
func RegisterSubCode(namespace string, subCode int) {
	subCodes[namespace] = SubCode(subCode)
}

// GetCode retrieves a code from namespace or Code(0)
func GetCode(namespace string) Code {
	if value, ok := codes[namespace]; ok {
		return value
	}
	return Code(0)
}

// GetSubCode retrieves a subcode from namespace or Code(0)
func GetSubCode(namespace string) SubCode {
	if value, ok := subCodes[namespace]; ok {
		return value
	}
	return SubCode(0)
}
