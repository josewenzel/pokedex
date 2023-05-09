package router

import (
	"net/http"
	"pokedex/src/api/handler"
)

func RouteStatusRequest(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		handler.HandleGetStatus(responseWriter, request)
	default:
		handler.StatusRespondWith(http.StatusNotImplemented, responseWriter)
	}
}
