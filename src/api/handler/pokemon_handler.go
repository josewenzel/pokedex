package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokedex/src/api/security"
	"pokedex/src/domain/service"
	"strconv"
)

func HandleGetPokemons(responseWriter http.ResponseWriter, request *http.Request, service service.PokemonService) {
	if !security.AuthorizeRequest(request).Valid {
		StatusRespondWith(http.StatusUnauthorized, responseWriter)
		return
	}

	id := request.URL.Query().Get("id")
	pokemonId, _ := strconv.Atoi(id)
	pokemon := service.GetPokemon(pokemonId)

	PokemonRespondWith(pokemon, http.StatusOK, responseWriter)
}

func PokemonRespondWith(pokemon service.Pokemon, statusCode int, responseWriter http.ResponseWriter) {
	responseWriter.WriteHeader(statusCode)
	jsonResponse, _ := json.Marshal(pokemon)
	fmt.Fprintf(responseWriter, string(jsonResponse))
}
