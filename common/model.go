package common

import "fmt"

// RequestParams holds a request parameters
type RequestParams struct {
	Int1  int    `json:"int1"`
	Int2  int    `json:"int2"`
	Limit int    `json:"limit"`
	Str1  string `json:"str1"`
	Str2  string `json:"str2"`
}

// GetID returns a unique request idenfier
func (rp RequestParams) GetID() string {
	return fmt.Sprintf("%v-%v-%v-%v-%v", rp.Int1, rp.Int2, rp.Limit, rp.Str1, rp.Str2)
}

// IsValid is the fizz buzz request valid ?
func (rp RequestParams) IsValid() bool {
	return rp.Limit >= 0
}

// RequestStats list a request params and its statistics (in our case hits)
type RequestStats struct {
	Request *RequestParams `json:"params"`
	Hits    int            `json:"hits"`
}
