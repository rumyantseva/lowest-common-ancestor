package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response represents response body.
type Response struct {
	Error *RespError  `json:"error,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

// RespError represents part of response body responsible for error displaying.
type RespError struct {
	Message string `json:"message"`
}

// WriteResponseError writes a response in case of error.
func WriteResponseError(w http.ResponseWriter, message string, code int) {
	resp := &Response{
		Error: &RespError{Message: message},
		Data:  nil,
	}
	writeResponse(w, resp, code)
}

// WriteResponseSuccess writes a response in case of success.
func WriteResponseSuccess(w http.ResponseWriter, data interface{}, code int) {
	resp := &Response{
		Error: nil,
		Data:  data,
	}
	writeResponse(w, resp, code)
}

func writeResponse(w http.ResponseWriter, resp *Response, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		log.Printf("Couldn't encode response %+v to JSON response body.", resp)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
