package web

import (
	"encoding/json"
	"log"
	"net/http"
)

// RespondWithError sends an error response with a custom message and logs the actual error.
func RespondWithError(w http.ResponseWriter, statusCode int, messageToSend string, err error) {
	log.Printf("err: %v", err)

	type ResponseStruct struct {
		Message string `json:"msg"`
	}

	res := ResponseStruct{
		Message: messageToSend,
	}

	data, er := json.Marshal(res)
	if er != nil {
		log.Printf("json.Marshal error: %v", er)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(messageToSend))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

// RespondWithJSON sends a JSON response with the given status code and data.
func RespondWithJSON(w http.ResponseWriter, statusCode int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("json.Marshal error in RespondWithJSON: %v", err)
		RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
