package handler

import (
	"fmt"
	"net/http"
	"net/url"
)

func StatusHandler(responseWriter http.ResponseWriter, request *http.Request) {
	_, err := url.Parse(request.URL.RawQuery)

	if err != nil {
		responseWriter.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(responseWriter, "invalid request")
		return
	}

	responseWriter.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(responseWriter, "{\"status\": \"OK\"}")
}
