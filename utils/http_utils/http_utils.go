package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/Seanlinsanity/golang-microservices-practice/src/api/utils/errors"
)

func RespondJson(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondErr(w http.ResponseWriter, err errors.ApiError) {
	RespondJson(w, err.Status(), err)
}
