package pkg

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, result int) {
	response := map[string]int{"result": result}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
