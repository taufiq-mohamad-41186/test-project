package restserver

import (
	"encoding/json"
	"net/http"
)

type (
	Meta struct {
		Count   int    `json:"count,omitempty"`
		Status  int    `json:"status"`
		Message string `json:"message"`
		Type    string `json:"type,omitempty"`
	}
	Response struct {
		Meta  *Meta       `json:"meta,omitempty"`
		Data  interface{} `json:"data,omitempty"`
		Error string      `json:"error,omitempty"`
	}
)

func Success(meta *Meta, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(meta.Status)
	err := json.NewEncoder(w).Encode(Response{meta, data, ""})
	if err != nil {
		println(err.Error())
		return
	}
}

func Failed(error string, statusCode int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	resp := Response{
		Meta: &Meta{
			Count:  0,
			Status: statusCode,
		},
		Data:  nil,
		Error: error,
	}

	err := json.NewEncoder(w).Encode(resp)

	if err != nil {
		println(err.Error())
		return
	}
}
