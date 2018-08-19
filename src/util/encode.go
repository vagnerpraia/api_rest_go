package util

import (
	"encoding/json"
	"net/http"

	"github.com/api_rest_go/src/model"
)

func EncodeResponseJson(w http.ResponseWriter, response model.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)

	enc := json.NewEncoder(w)
	enc.Encode(response)
}
