package main

import (
	"encoding/json"
	"net/http"
)

/*
*
@desc convert data to JSON and write to http response
@var w http.ResponseWriter write data to http response
@var data any data payload want to return to client
*/
func writeJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data) // convert data to JSON and write to v
}

/*
@desc Read and decode JSON data from the HTTP request body (send from client) into the given struct or map.
@var w http.ResponseWriter HTTP response writer (used to enforce request body size limits).
@var r *http.Request HTTP request containing the JSON payload.
@var data any: A pointer to the struct or map where the decoded JSON data will be stored.
@return error: Returns an error if the JSON is invalid, exceeds the size limit, or contains unknown fields.
*/
func readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_576 // 1MB
	// limit maximum bytes read from request body
	// if error, return request body too large
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	decoder := json.NewDecoder(r.Body)
	// not allowed unknown fields in JSON
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

/*
@desc Write a JSON error response to the HTTP client.
@var w http.ResponseWriter HTTP response writer to send the JSON response.
@var status int: HTTP status code for the error response.
@var message string: Error message to be included in the response.
@return error: Returns an error if writing the JSON response fails.
*/
func writeJSONError(w http.ResponseWriter, status int, message string) error {
	type envelop struct {
		Error string `json:"error"`
	}

	return writeJSON(w, status, &envelop{Error: message})
}
