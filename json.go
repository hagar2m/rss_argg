package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(httpWiter http.ResponseWriter, code int, msg string)  {
	if code > 499 {
		log.Println("Responging with 5xx error", msg)
	}
	type errResponse struct {
		 Error string `json:"error"`
	}

	responseWithJSON(httpWiter, code, errResponse{
		Error: msg,
	})
}

func responseWithJSON(httpWiter http.ResponseWriter, code int, payload interface{})  {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling json response: %v", payload)
		httpWiter.WriteHeader(500)
		return
	}
	httpWiter.Header().Add("Content-Type", "application/json")
	httpWiter.WriteHeader(code)
	httpWiter.Write(data)

}