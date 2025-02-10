package v1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mattcardoso/foxbit-tech-test/pkg"
)

func TestSumHandleFuncOkPath(t *testing.T) {
	tt := []struct {
		t1       string
		t2       string
		expected map[string]int
	}{
		{t1: "5", t2: "5", expected: map[string]int{"result": 10}},
		{t1: "5", t2: "0", expected: map[string]int{"result": 5}},
	}
	for _, tc := range tt {
		path := fmt.Sprintf("/api/sum?term_one=%s&term_two=%s", tc.t1, tc.t2)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		mux := http.NewServeMux()
		mux.HandleFunc("GET /api/sum", Sum)
		mux.ServeHTTP(rr, req)

		var response map[string]int
		json.Unmarshal(rr.Body.Bytes(), &response)
		if response["result"] != tc.expected["result"] {
			t.Errorf("wrong response for params (t1,t2) (%s,%s), want: %v got: %v", tc.t1, tc.t2, tc.expected, response)
		}
	}
}

func TestSumHandleFuncWrongPathEmptyValues(t *testing.T) {
	tt := []struct {
		t1            string
		t2            string
		expectedCode  int
		expectedError string
	}{
		{t1: "", t2: "5", expectedCode: http.StatusBadRequest, expectedError: pkg.ErrEmptyTerms.Error() + "\n"},
		{t1: "5", t2: "", expectedCode: http.StatusBadRequest, expectedError: pkg.ErrEmptyTerms.Error() + "\n"},
	}
	for _, tc := range tt {
		path := fmt.Sprintf("/api/sum?term_one=%s&term_two=%s", tc.t1, tc.t2)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		mux := http.NewServeMux()
		mux.HandleFunc("GET /api/sum", Sum)
		mux.ServeHTTP(rr, req)

		var response string = rr.Body.String()
		if response != tc.expectedError {
			t.Errorf("wrong response for params (t1,t2) (%s,%s), want: '%s' got: '%s'", tc.t1, tc.t2, tc.expectedError, response)
		}
		if rr.Code != tc.expectedCode {
			t.Errorf("wrong reponse code want: %v got: %v", tc.expectedCode, rr.Code)
		}
	}
}

func TestSumHandleFuncWrongPathNonNumberValues(t *testing.T) {
	tt := []struct {
		t1            string
		t2            string
		expectedCode  int
		expectedError string
	}{
		{t1: "12a", t2: "5", expectedCode: http.StatusBadRequest, expectedError: pkg.ErrTermsNotNumber.Error() + "\n"},
		{t1: "5", t2: "12a", expectedCode: http.StatusBadRequest, expectedError: pkg.ErrTermsNotNumber.Error() + "\n"},
	}
	for _, tc := range tt {
		path := fmt.Sprintf("/api/sum?term_one=%s&term_two=%s", tc.t1, tc.t2)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()

		mux := http.NewServeMux()
		mux.HandleFunc("GET /api/sum", Sum)
		mux.ServeHTTP(rr, req)

		var response string = rr.Body.String()
		if response != tc.expectedError {
			t.Errorf("wrong response for params (t1,t2) (%s,%s), want: '%s' got: '%s'", tc.t1, tc.t2, tc.expectedError, response)
		}
		if rr.Code != tc.expectedCode {
			t.Errorf("wrong reponse code want: %v got: %v", tc.expectedCode, rr.Code)
		}
	}
}

// want: term_one and term_two must be set and not empty got: term_one and term_two must be set and not empty
