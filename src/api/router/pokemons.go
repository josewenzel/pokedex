package router

import (
	"net/http"
	"pokedex/src/api/handler"
)

func RoutePokemonsRequest(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		handler.HandleGetPokemons(responseWriter, request)
	}
}
