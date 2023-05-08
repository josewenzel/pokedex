package handler

import (
	"net/http"
)

func HandleGetPokemons(responseWriter http.ResponseWriter, request *http.Request) {
	StatusRespondWith(http.StatusUnauthorized, responseWriter)
	return
}
