package response

import (
	"encoding/json"
	"net/http"
)

func DataResponse(w http.ResponseWriter, data interface{}) {
	response := Response{
		Code: http.StatusOK,
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to serialize response", http.StatusInternalServerError)
	}
}

func ErrResponse(w http.ResponseWriter, err error, code int) *Response {
	response := Response{
		Code:  code,
		Error: err.Error(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(response)

	return &response
}
