package errx

// Error describes an error with more information
type Error struct {
	Code    Code    `json:"code"`
	SubCode SubCode `json:"sub_code"`
	Message string  `json:"message"`
	Raw     error   `json:"-"`
}
