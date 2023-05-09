package router

import (
	"net/http"
	"pokedex/src/api/handler"
	"pokedex/src/domain/service"
)

type PokemonRouteHandler struct {
	PokemonService service.PokemonService
}

func (r PokemonRouteHandler) RoutePokemonsRequest(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		handler.HandleGetPokemons(responseWriter, request, r.PokemonService)
	}
}
