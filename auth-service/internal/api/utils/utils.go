/**
 * @author Mory Keita on 1/20/20
 */
package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
		)

// Represents error structure for generic error.
// use to create standard API response template.
type GenericError struct {
	code int `json:"code"`
	Error string `json:"error"`
	Data interface{} `json:"data,omitempty"`
} 

// returns a JSON error message and HTTP status code.
func WriteError(w http.ResponseWriter , code int , message string,data interface{})  {
	response := GenericError{
		code:  code,
		Error: message,
		Data:  data,
	}

	WriteJSON(w,code,response)
}

// returns a JSON data and HTTP status code.
func WriteJSON( w http.ResponseWriter , code int,data interface{})  {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(data); err != nil{
		log.WithError(err).Warn("Error writing response.")
	}
}
