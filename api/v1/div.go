package v1

import (
	"net/http"

	"github.com/mattcardoso/foxbit-tech-test/pkg"
)

func Div(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	if err := pkg.ValidateParams(q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t1, t2, err := pkg.ParseParams(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if t2 == 0 {
		http.Error(w, pkg.ErrDivideByZero.Error(), http.StatusBadRequest)
		return
	}
	pkg.JsonResponse(w, t1/t2)
}
