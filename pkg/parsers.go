package pkg

import (
	"net/url"
	"strconv"
)

func ValidateParams(q url.Values) error {
	term_one, term_two := q.Get("term_one"), q.Get("term_two")
	if term_one == "" || term_two == "" {
		return ErrEmptyTerms
	}
	return nil
}

func ParseParams(q url.Values) (int, int, error) {
	term_one, term_two := q.Get("term_one"), q.Get("term_two")
	int_one, err := strconv.Atoi(term_one)
	if err != nil {
		return 0, 0, ErrTermsNotNumber
	}
	int_two, err := strconv.Atoi(term_two)
	if err != nil {
		return 0, 0, ErrTermsNotNumber
	}
	return int_one, int_two, nil
}
