package pkg

import "fmt"

var (
	ErrEmptyTerms     = fmt.Errorf("term_one and term_two must be set and not empty")
	ErrTermsNotNumber = fmt.Errorf("all terms must be intergers")
	ErrDivideByZero   = fmt.Errorf("divisor must not be 0")
)
