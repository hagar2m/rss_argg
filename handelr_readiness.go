package main

import "net/http"

func handlerReadiness(httpWiter http.ResponseWriter, request *http.Request)  {
	responseWithJSON(httpWiter, 200, struct{}{})
}