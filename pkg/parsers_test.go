package pkg

import (
	"net/url"
	"testing"
)

// Test ValidateParams
func TestValidateParams(t *testing.T) {
	tests := []struct {
		name     string
		query    url.Values
		expected error
	}{
		{"Valid params", url.Values{"term_one": {"5"}, "term_two": {"10"}}, nil},
		{"Missing term_one", url.Values{"term_one": {"5"}, "term_two": {""}}, ErrEmptyTerms},
		{"Missing term_two", url.Values{"term_one": {"5"}}, ErrEmptyTerms},
		{"Both terms missing", url.Values{}, ErrEmptyTerms},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateParams(tt.query)
			if err != tt.expected {
				t.Errorf("Expected error: %v, got: %v", tt.expected, err)
			}
		})
	}
}

// Test ParseParams
func TestParseParams(t *testing.T) {
	tests := []struct {
		name        string
		query       url.Values
		expectedOne int
		expectedTwo int
		expectedErr error
	}{
		{"Valid numbers", url.Values{"term_one": {"5"}, "term_two": {"10"}}, 5, 10, nil},
		{"term_one not valid", url.Values{"term_one": {"abc"}, "term_two": {"10"}}, 0, 0, ErrTermsNotNumber},
		{"term_two not valid", url.Values{"term_one": {"5"}, "term_two": {"xyz"}}, 0, 0, ErrTermsNotNumber},
		{"Both not valid", url.Values{"term_one": {"abc"}, "term_two": {"xyz"}}, 0, 0, ErrTermsNotNumber},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intOne, intTwo, err := ParseParams(tt.query)
			if intOne != tt.expectedOne || intTwo != tt.expectedTwo || err != tt.expectedErr {
				t.Errorf("Expected (%d, %d, %v), got (%d, %d, %v)",
					tt.expectedOne, tt.expectedTwo, tt.expectedErr,
					intOne, intTwo, err)
			}
		})
	}
}
