package utils

import (
    "net/http"
    "encoding/json"
)

type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
}

func JSONResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    
    response := Response{
        Status:  http.StatusText(statusCode),
        Message: message,
        Data:    data,
    }
    
    json.NewEncoder(w).Encode(response)
}