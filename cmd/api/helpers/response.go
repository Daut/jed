package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/daut/jed/internal/utils"
)

type Response struct {
	Logger *utils.Logger
}

func NewResponse(logger *utils.Logger) *Response {
	return &Response{
		Logger: logger,
	}
}

type errorResponse struct {
	Message string `json:"message"`
}

// Returns a 500 Internal Server Error response to the client
func (h *Response) ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	h.Logger.Error.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Returns a 400 Bad Request response to the client
func (h *Response) ClientError(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&errorResponse{Message: msg})
}

// Returns a 404 Not Found response to the client
func (h *Response) NotFound(w http.ResponseWriter) {
	msg := "the requested resource could not be found"
	h.ClientError(w, msg, http.StatusNotFound)
}

// Writes a JSON response to the client
func (h *Response) WriteJSON(w http.ResponseWriter, status int, data any, headers http.Header) {
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
