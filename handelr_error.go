package main

import "net/http"

func handlerError(httpWiter http.ResponseWriter, request *http.Request)  {
	responseWithError(httpWiter, 400, "Something went wrong")
}