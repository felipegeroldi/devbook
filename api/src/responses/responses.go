package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON return an response in JSON to request
func JSON(rw http.ResponseWriter, statusCode int, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(rw).Encode(data); err != nil {
			log.Fatal(err)
		}
	}
}

// Error returns an error in JSON format to request
func Error(rw http.ResponseWriter, statusCode int, err error) {
	JSON(rw, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
